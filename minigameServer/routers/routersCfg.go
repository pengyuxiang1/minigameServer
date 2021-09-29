package routers

import (
	"github.com/beego/beego/v2/server/web"
	"minigameServer/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/login", &controllers.LoginController{})
	web.Router("/register", &controllers.RegController{})
	web.Router("/logout",&controllers.LoginController{},"*:Logout")

}