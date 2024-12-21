package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}
