package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model

	Quote string `gorm:"not null" json:"quote"`
	Author string `gorm:"not null" json:"author"`
	Source string `gorm:"not null" json:"source"`
}