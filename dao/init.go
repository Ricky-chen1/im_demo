package dao

import (
	"imgo/model"
	"imgo/pkg/util"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func MySQLInit(path string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default.LogMode(logger.Warn)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       path,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
		DontSupportRenameIndex:    true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //启用单数表名
		},
	})

	if err != nil {
		util.LogInstance.Error("connect to Database failed", err)
		panic(err)
	}

	DB = db

	migration()
}

// 自动迁移
func migration() {
	DB.AutoMigrate(&model.User{}, &model.Msg{}, &model.Friend{}, &model.Request{}, &model.Group{}, &model.Group_member{})
}
