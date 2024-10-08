package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	UUID, _ := uuid.NewRandom()
	u.ID = UUID.String()
	return
}
