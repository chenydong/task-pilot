package main

import (
	"chat/router"
	"fmt"
	"task_pilot/config"
)

func main() {
	ctx, err := config.Init()
	if err != nil {
		panic("配置倒入错误")
	}
	r := router.ResisterRouter(ctx)
	err = r.Run(ctx.Config.Port)
	if err != nil {
		fmt.Println(err)
		panic("启动错误")
	}
}
