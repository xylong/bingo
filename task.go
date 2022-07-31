package bingo

import "sync"

var (
	once      sync.Once
	taskQueue chan *TaskExecutor
)

func init() {
	tasks := getTaskQueue()

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

func getTaskQueue() chan *TaskExecutor {
	once.Do(func() {
		taskQueue = make(chan *TaskExecutor)
	})

	return taskQueue
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
			getTaskQueue() <- NewTaskExecutor(f, param, callback)
		}()
	}
}
