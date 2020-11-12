package task

import (
	"encoding/json"
	"image/png"
	"io"
	"os"

	"fabu.dev/api/pkg/short"

	"fabu.dev/api/pkg/api/global"
	"fabu.dev/api/pkg/parser"

	"fabu.dev/api/pkg/config"

	"fabu.dev/api/pkg/constant"

	"fabu.dev/api/pkg/db"

	"fabu.dev/api/service"
	"github.com/sirupsen/logrus"
)

// 组合分片上传的文件数据
type CombineApp struct {
	BaseTask
	AppSaveRootPath string
	FileBuffer      map[string]map[int]*service.UploadInfo
}

func NewCombineApp() *CombineApp {
	return &CombineApp{
		AppSaveRootPath: config.Conf.System.AppSaveRootPath,
		FileBuffer:      make(map[string]map[int]*service.UploadInfo, 8),
	}
}

// 读取channel
func (t *CombineApp) Execute() {
	select {
	case data := <-service.FileChannel:
		err := t.Combine(data)
		if err != nil { // TODO 如果合并过程出现错误，需要处理
			logrus.Error(err)
		}

		return
	default:
		return
	}
}

// 写数据文件中
// channel中的数据有多个文件的数据，要根据文件标识来判断是创建文件还是续写文件，顺序是否正确
func (t *CombineApp) Combine(data *service.UploadInfo) error {
	// 从Redis中获取文件上传进度
	progress, _ := db.Redis.HGet(constant.AppUploadProgress, data.Params.Identifier).Int()

	// todo 验证是不是下一个分片，如果不是要仍会channel，或者暂存到sync.map
	// 如果data不是下一个需要合并的分片，暂存到map中。
	if data.Params.ChunkNumber > progress+1 {
		t.BufferFragment(data)
		return nil
	}

	if err := t.Save(data); err != nil {
		logrus.Error("save err", err)
		return err
	}

	// 从map中查找下次合并的分片
	if err := t.ReadBuffer(data.Params.Identifier, data.Params.ChunkNumber+1); err != nil {
		return err
	}

	return nil
}

// 从buffer内读取分片，进行保存
func (t *CombineApp) ReadBuffer(identifier string, nextChunk int) error {
	for {
		if _, ok := t.FileBuffer[identifier]; !ok {
			break
		}

		if _, ok := t.FileBuffer[identifier][nextChunk]; !ok {
			break
		}

		if err := t.Save(t.FileBuffer[identifier][nextChunk]); err != nil {
			return err
		}
		nextChunk++
	}

	return nil
}

// 缓存分片片段
func (t *CombineApp) BufferFragment(data *service.UploadInfo) {
	if bufferApp, ok := t.FileBuffer[data.Params.Identifier]; ok {
		bufferApp[data.Params.ChunkNumber] = data
		return
	}

	t.FileBuffer[data.Params.Identifier] = make(map[int]*service.UploadInfo, 128)
	t.FileBuffer[data.Params.Identifier][data.Params.ChunkNumber] = data

	return
}

// 保存文件到硬盘上
func (t *CombineApp) Save(data *service.UploadInfo) error {
	filename := t.AppSaveRootPath + data.Params.Identifier + data.Params.Filename
	out, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // TODO 这里感觉还可以优化，避免多次打开关闭文件
	if err != nil {
		return err
	}
	defer out.Close()

	src, err := data.Params.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	// 通过redis记录每个app上传进度
	if err := db.Redis.HSet(constant.AppUploadProgress, data.Params.Identifier, data.Params.ChunkNumber).Err(); err != nil {
		return err
	}

	// 如果最后一个分片合并完成，那么解析这个文件将数据保存到redis，删除FileBuffer内的数据
	if data.Params.ChunkNumber == data.Params.ChunkTotal {
		if err := t.CombineFinished(filename, data.Params.Identifier); err != nil {
			return nil
		}
	}

	return nil
}

// 合并完成，那么解析这个文件将数据保存到redis，删除FileBuffer内的数据
func (t *CombineApp) CombineFinished(filename, identifier string) error {
	apk, err := parser.NewAppParser(filename)
	if err != nil {
		return err
	}

	iconName := t.AppSaveRootPath + identifier + ".png"
	out, err := os.Create(iconName)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, apk.Icon)
	if err != nil {
		return err
	}

	shortKey, err := short.NewPool().GetShortKey()
	if err != nil {

	}
	appInfo := &global.AppInfo{
		Name:       apk.Name,
		BundleId:   apk.BundleId,
		Version:    apk.Version,
		Build:      apk.Build,
		Icon:       iconName,
		Size:       apk.Size,
		Identifier: identifier,
		Platform:   apk.Platform,
		ShortKey:   shortKey,
	}
	data, err := json.Marshal(appInfo)
	if err != nil {
		return err
	}

	if err := db.Redis.HSet(constant.AppFileInfo, identifier, data).Err(); err != nil { // TODO 这里是否要改用string？
		logrus.Error("save file info err", err)
		return err
	}

	delete(t.FileBuffer, identifier) // 完成后删除buffer内的缓存

	db.Redis.HDel(constant.AppUploadProgress, identifier)

	return nil
}
