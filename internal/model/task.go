package model

import (
	"database/sql"
	"time"

	"github.com/caclm10/simpletodo-api/internal/model/response"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	TodoID    uint
	Content   string
	Sequence  uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Todo      *Todo
	Steps     []Step
}

func (t Task) ToResponse() response.TaskResponse {
	return response.TaskResponse{
		ID:       t.ID,
		Content:  t.Content,
		Sequence: t.Sequence,
	}
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	var maxseq sql.NullInt64

	if err := tx.Model(&Task{}).Select("MAX(sequence)").Where("todo_id = ?", t.TodoID).Row().Scan(&maxseq); err != nil {
		return err
	}

	if maxseq.Valid {
		t.Sequence = uint(maxseq.Int64) + 1
	} else {
		t.Sequence = 0
	}

	return nil
}
