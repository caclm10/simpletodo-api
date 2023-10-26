package model

import (
	"database/sql"
	"time"

	"github.com/caclm10/simpletodo-api/internal/model/response"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string
	Sequence  uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	User      *User
	Tasks     []Task
}

func (t Todo) ToResponse() response.TodoResponse {
	return response.TodoResponse{
		ID:       t.ID,
		Name:     t.Name,
		Sequence: t.Sequence,
	}
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	var maxseq sql.NullInt64

	if err := tx.Model(&Todo{}).Select("MAX(sequence)").Where("user_id = ?", t.UserID).Row().Scan(&maxseq); err != nil {
		return err
	}

	if maxseq.Valid {
		t.Sequence = uint(maxseq.Int64) + 1
	} else {
		t.Sequence = 0
	}

	return nil
}
