package model

import "gorm.io/gorm"

type Group_member struct {
	gorm.Model
	Uid      uint   `gorm:"not null"`
	Gid      uint   `gorm:"not null"`
	NickName string `gorm:"indx"`       //群昵称
	Mute     int    `gorm:"default:-1"` //是否禁言 -1表示未禁言,1表示禁言
}
