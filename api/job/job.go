package job

import (
	"time"

	"fabu.dev/api/pkg/utils"
)

type Command struct {
	taskList []Task
}

func NewCommand() *Command {
	return &Command{
		taskList: TaskList,
	}
}

func (i *Command) Process() {
	if len(i.taskList) == 0 {
		return
	}

	for _, creator := range i.taskList {
		// 开始执行各个任务
		go i.execute(creator)
	}
}

func (i *Command) execute(creator Task) {
	defer utils.HandlePanic()

	for {
		creator.Execute()

		//s := reflect.TypeOf(creator).String()
		//fmt.Printf("%s 执行一次\n", s)
		time.Sleep(creator.GetDelayTime())
	}
}

func (i *Command) SetCreator(creator Task) {
	i.taskList = append(i.taskList, creator)
}
