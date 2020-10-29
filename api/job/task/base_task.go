package task

import "time"

type BaseTask struct {
}

func NewBaseTask() *BaseTask {
	return &BaseTask{}
}

func (t BaseTask) GetDelayTime() time.Duration {
	return time.Millisecond * 500
}
