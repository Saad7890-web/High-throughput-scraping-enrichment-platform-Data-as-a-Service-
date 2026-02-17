package job

import (
	"time"

	"github.com/Saad7890-web/scrapper-platform/internal/worker"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
	pool *worker.Pool
}

func NewService(repo *Repository, pool *worker.Pool) *Service{
	return &Service{repo: repo, pool: pool}
}

func(s *Service)Create(url string)(*Job, error){
	j := &Job{
		ID: uuid.New().String(),
		URL: url,
		Status: StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(j); err != nil {
		return nil, err
	}

	s.pool.Submit(worker.Task{
		ID: j.ID,
		URL: j.URL,
	})

	return j, nil
}