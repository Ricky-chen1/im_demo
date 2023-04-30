package model

import (
	"gorm.io/gorm"
)

type Msg struct {
	gorm.Model
	From        string `gorm:"not null"`
	To          string //id 可能是群聊id或者用户id
	MessageType uint   `gorm:"default:1"` // 1为单聊,2为群聊
	Content     string `gorm:"not null"`
	ExpireTime  int64  `gorm:"index"`
}
