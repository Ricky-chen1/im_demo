package main

import (
	"imgo/conf"
	"imgo/pkg/util"
	"imgo/router"
	"imgo/service/ws"
)

func main() {
	//路由注册
	r := router.Init()

	//开启监听websocket连接
	go ws.Manager.Start()

	//开启服务
	err := r.Run(conf.HttpPort)
	if err != nil {
		util.LogInstance.Fatal("ListenAndServe: ", err)
	}
}
