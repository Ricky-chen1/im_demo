package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var LogInstance *logrus.Logger

func init() {

	//创建对象
	logger := logrus.New()

	//设置输出
	logger.Out = os.Stdout

	//设置日志等级
	logger.SetLevel(logrus.InfoLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	LogInstance = logger

}
