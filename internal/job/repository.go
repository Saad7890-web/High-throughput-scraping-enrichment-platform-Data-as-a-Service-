package job

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(j *Job) error {
	_, err := r.db.Exec(`
		INSERT INTO jobs (id, url, status, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5)
	`,
		j.ID, j.URL, j.Status, j.CreatedAt, j.UpdatedAt,
	)
	return err
}

func (r *Repository) UpdateStatus(id string, status Status, errMsg *string) error {
	_, err := r.db.Exec(`
		UPDATE jobs
		SET status=$2, error=$3, updated_at=$4
		WHERE id=$1
	`, id, status, errMsg, time.Now())

	return err
}
