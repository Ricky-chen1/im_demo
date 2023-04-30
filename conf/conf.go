package conf

import (
	"imgo/dao"
	"imgo/pkg/util"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	appMode  string
	HttpPort string

	db         string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
)

func init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		util.LogInstance.Info("load conf.ini failed", err)
	}

	loadServer(file)
	loadMySQL(file)
	path := strings.Join([]string{dbUser, ":", dbPassword, "@tcp(", dbHost, ":", dbPort, ")/",
		dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	dao.MySQLInit(path)
}

func loadServer(file *ini.File) {
	s := file.Section("service")
	appMode = s.Key("appMode").String()
	HttpPort = s.Key("httpPort").String()
}

func loadMySQL(file *ini.File) {
	s := file.Section("mysql")
	dbHost = s.Key("dbHost").String()
	dbPort = s.Key("dbPort").String()
	dbUser = s.Key("dbUser").String()
	dbPassword = s.Key("dbPassword").String()
	dbName = s.Key("dbName").String()
}
