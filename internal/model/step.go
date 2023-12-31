package model

import (
	"database/sql"
	"time"

	"github.com/caclm10/simpletodo-api/internal/model/response"
	"gorm.io/gorm"
)

type Step struct {
	ID        uint `gorm:"primaryKey"`
	TaskID    uint
	Content   string
	Sequence  uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Task      *Task
}

func (s Step) ToResponse() response.StepResponse {
	return response.StepResponse{
		ID:       s.ID,
		Content:  s.Content,
		Sequence: s.Sequence,
	}
}

func (t *Step) BeforeCreate(tx *gorm.DB) error {
	var maxseq sql.NullInt64

	if err := tx.Model(&Step{}).Select("MAX(sequence)").Where("task_id = ?", t.TaskID).Row().Scan(&maxseq); err != nil {
		return err
	}

	if maxseq.Valid {
		t.Sequence = uint(maxseq.Int64) + 1
	} else {
		t.Sequence = 0
	}

	return nil
}
