package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UUID      uuid.UUID      `json:"uuid" gorm:"type:char(36);uniqueIndex;not null"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (c *Categories) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New()
	return
}
