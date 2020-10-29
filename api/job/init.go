package job

import (
	"time"

	"fabu.dev/api/job/task"
)

var TaskList = make([]Task, 0, 8)

type Task interface {
	Execute()
	GetDelayTime() time.Duration
}

// 此处添加要执行的job
func init() {
	TaskList = append(TaskList, task.NewCombineApp())
}
