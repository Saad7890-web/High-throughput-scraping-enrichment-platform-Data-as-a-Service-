package job

import "log"

type Executor struct {
	repo *Repository
}

func NewExecutor(repo *Repository) *Executor {
	return &Executor{repo: repo}
}

func (e *Executor) OnStart(id string) {
	e.repo.UpdateStatus(id, StatusRunning, nil)
}

func (e *Executor) OnSuccess(id string) {
	e.repo.UpdateStatus(id, StatusCompleted, nil)
}

func (e *Executor) OnFailure(id string, err error) {
	msg := err.Error()
	e.repo.UpdateStatus(id, StatusFailed, &msg)
	log.Println("job failed:", msg)
}
