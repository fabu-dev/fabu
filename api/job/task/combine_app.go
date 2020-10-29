package task

import (
	"io"
	"os"
	"path/filepath"

	"fabu.dev/api/service"
	"github.com/sirupsen/logrus"
)

// 组合分片上传的文件数据
type CombineApp struct {
	BaseTask
}

func NewCombineApp() *CombineApp {
	return &CombineApp{}
}

// 读取channel
// channel中的数据有多个文件的数据，要根据文件标识来判断是创建文件还是续写文件，顺序是否正确
func (t *CombineApp) Execute() {
	select {
	case data := <-service.FileChannel:
		t.Save(data)
		return
	default:
		return
	}
}

// 写数据文件中
func (t *CombineApp) Save(data *service.UploadInfo) error {
	logrus.Info("get save date is :", data)

	filename := filepath.Base(data.Params.File.Filename)

	src, err := data.Params.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		logrus.Info("write file err ", err)
	}
	return err
}
