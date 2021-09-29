package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
)


type MainController struct {
	web.Controller
}

var globalSessions *session.Manager
func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}
type JSONS struct {
	//必须的大写开头
	Code string
	Msg  string
	User []string `json:"user_info"`//key重命名,最外面是反引号
}

func (c *MainController) Get() {
	data := &JSONS{"100", "成功连接",
		[]string{"minigame","music"}}
	c.Data["json"] = data
	c.ServeJSON()
}