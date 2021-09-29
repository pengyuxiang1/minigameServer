package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "minigameServer/routers"
)

func main() {
	web.Run()
}