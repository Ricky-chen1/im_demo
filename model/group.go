package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string // 群名
	OwnerID uint   `gorm:"not null;index"` //群主id
}
