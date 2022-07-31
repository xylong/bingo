package bingo

import (
	"github.com/robfig/cron/v3"
	"sync"
)

var (
	taskOnce     sync.Once
	cronOnce     sync.Once
	taskList     chan *TaskExecutor
	cronTaskList *cron.Cron
)

func init() {
	tasks := getTaskList()

	go func() {
		for task := range tasks {
			exec(task)
		}
	}()
}

// exec 执行任务
func exec(executor *TaskExecutor) {
	go func() {
		defer func() {
			if executor.callback != nil {
				executor.callback()
			}
		}()

		executor.Do()
	}()
}

func getTaskList() chan *TaskExecutor {
	taskOnce.Do(func() {
		taskList = make(chan *TaskExecutor)
	})

	return taskList
}

// TaskFunc 协程任务方法
type TaskFunc func(param ...interface{})

// TaskExecutor 任务执行器
type TaskExecutor struct {
	f        TaskFunc
	p        []interface{}
	callback func()
}

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

// Do 执行任务
func (t *TaskExecutor) Do() {
	t.f(t.p...)
}

// Task 加入任务队列
func Task(f TaskFunc, callback func(), param ...any) {
	if f != nil {
		go func() {
			getTaskList() <- NewTaskExecutor(f, param, callback)
		}()
	}
}

func getCron() *cron.Cron {
	cronOnce.Do(func() {
		cronTaskList = cron.New(cron.WithSeconds())
	})

	return cronTaskList
}
