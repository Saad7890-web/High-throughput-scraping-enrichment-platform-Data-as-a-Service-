package worker

import (
	"log"
	"time"
)

type Executor interface {
	OnStart(id string)
	OnSuccess(id string)
	OnFailure(id string, err error)
}

type DefaultProcessor struct {
	exec Executor
}

func NewProcessor(exec Executor) *DefaultProcessor {
	return &DefaultProcessor{exec: exec}
}

func (p *DefaultProcessor) Process(id string, url string) {

	p.exec.OnStart(id)


	time.Sleep(2 * time.Second)

	log.Println("scraped:", url)

	p.exec.OnSuccess(id)
}
