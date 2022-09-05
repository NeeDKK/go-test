package main

import (
	"fmt"
	"time"
)

type Task struct {
	TaskId int
	t      func() error
}

// 创建任务
func NewTask(taskId int, f func() error) *Task {
	return &Task{
		TaskId: taskId,
		t:      f,
	}
}

// 执行任务
func (t *Task) Excute() {
	t.t()
}

type GoRoutinePool struct {
	WorkerNum int
	entryChan chan *Task //对外接收task的入口
	jobChan   chan *Task //对内准备就绪的队列
}

func NewWorkerPool(work int) *GoRoutinePool {
	return &GoRoutinePool{
		WorkerNum: work,
		entryChan: make(chan *Task),
		jobChan:   make(chan *Task),
	}
}
func (p *GoRoutinePool) Run() {
	for i := 0; i < p.WorkerNum; i++ {
		go p.work(i)
	}
	// 将写入的任务放入执行chan
	for job := range p.entryChan {
		p.jobChan <- job
	}

}
func (p *GoRoutinePool) work(workId int) {
	for task := range p.jobChan {
		task.t()
		fmt.Println("执行任务:", task.TaskId)
	}

}
func (p *GoRoutinePool) Close() {
	if _, ok := <-p.entryChan; ok {
		close(p.entryChan)
	}
	if _, ok := <-p.jobChan; ok {
		close(p.jobChan)
	}
}

func main() {
	p := NewWorkerPool(3)
	go func() {
		// 模拟任务写入
		for {
			time.Sleep(time.Second)
			p.entryChan <- NewTask(int(time.Now().Unix()), func() error {
				//fmt.Println("任务:", i, "", time.Now())
				return nil
			})
		}
	}()
	p.Run()
	defer p.Close()
}
