package model

import (
	"time"

	"gorm.io/gorm"
)

type Mail struct {
	Id       uint      `gorm:"column:id" json:"id"`
	Action    string    `gorm:"column:action" json:"action"`
	Receiver string    `gorm:"column:receiver" json:"receiver"`
	CreateAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *Mail) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	m.CreateAt = now
	return nil
}

func (m *Mail) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateAt = time.Now()
	return nil
}