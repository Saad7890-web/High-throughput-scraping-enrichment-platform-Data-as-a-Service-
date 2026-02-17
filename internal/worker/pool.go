package worker

import "log"


type JobProcessor interface{
	Process(jobID string, url string)
}

type Pool struct{
	workers int
	queue chan Task
}

type Task struct {
	ID string
	URL string
}

func NewPool(workers int) *Pool{
	return  &Pool{
		workers: workers,
		queue: make(chan Task, 1000),
	}
}

func (p *Pool) Start(processor JobProcessor){
	for i:= 0;i < p.workers; i++ {
		go func(id int){
			log.Println("Worker started", id)
			for task := range p.queue {
				processor.Process(task.ID, task.URL)
			}
		}(i)
	}
}

func (p *Pool) Submit(task Task){
	p.queue <- task
}