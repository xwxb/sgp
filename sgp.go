package sgp

import (
	"sync"
)

type Task struct {
	f func()
}

func NewTask(f func()) *Task {
	return &Task{f}
}

type Pool struct {
	taskChan  chan *Task
	wokerSize int
	wg        sync.WaitGroup
}

func NewPool(limit int) *Pool {
	return &Pool{
		taskChan:  make(chan *Task),
		wokerSize: limit,
	}
}

func (p *Pool) AddTask(task *Task) {
	p.taskChan <- task
}

func (p *Pool) Run() {
	for i := 0; i < p.wokerSize; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.taskChan {
				task.f()
			}
		}()
	}
}

func (p *Pool) Close() {
	close(p.taskChan)
	p.wg.Wait()
}
