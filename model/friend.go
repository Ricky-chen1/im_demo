package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	OwnerID    uint  `gorm:"index"` //该好友拥有者id
	FriendID   uint  `gorm:"uid"`   //好友id
	CreateTime int64 `gorm:"not null"`
}

//好友请求
type Request struct {
	gorm.Model
	ReqUid   uint `gorm:"not null"`
	To       uint `gorm:"not null"`
	IsAccept bool `gorm:"default:false"` //默认还未接受
}
