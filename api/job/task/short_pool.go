package task

import (
	"fmt"
	"time"

	"fabu.dev/api/pkg/config"
	"fabu.dev/api/pkg/short"

	"fabu.dev/api/service"
)

// 组合分片上传的文件数据
type ShortPool struct {
	BaseTask
	AppSaveRootPath string
	FileBuffer      map[string]map[int]*service.UploadInfo
}

func NewShortPool() *ShortPool {
	return &ShortPool{
		AppSaveRootPath: config.Conf.System.AppSaveRootPath,
		FileBuffer:      make(map[string]map[int]*service.UploadInfo, 8),
	}
}

// 持续监听pool，
func (t *ShortPool) Execute() {
	objPool := short.NewPool()

	err := objPool.Buffer()
	if err != nil {
		fmt.Printf("pool err %s\n", err.Error())
	}
}

// 1分钟执行1次
func (t ShortPool) GetDelayTime() time.Duration {
	return time.Second * 60
}
