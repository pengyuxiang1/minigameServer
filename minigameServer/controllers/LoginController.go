package controllers

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"io"
	"minigameServer/models"
)

type LoginController struct {
	web.Controller
}

//登陆处理
func (this *LoginController) Post() {
	this.Ctx.Request.ParseForm()
	accountId := this.Ctx.Request.Form.Get("account_id")
	password := this.Ctx.Request.Form.Get("password")
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()
	userInfo,err := models.GetUserInfoById(accountId)
	if err != nil {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "account_id error, Please to again",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	if userInfo.Password == newPass {
		//登录成功设置session
		sess,err := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)//全局共享session
		if err != nil {
			this.Data["json"] = models.RspComm{
				Code: 200,
				Msg:  "session设置失败，登录失败",
				Info: "",
			}
			this.ServeJSON()
			return
		}
		sess.Set(context.Background(),"LoginUserInfo",userInfo)
		this.Data["json"] = models.RspComm{
			Code: 100,
			Msg:  "登录成功",
			Info: sess.SessionID(context.Background()),
		}
		this.ServeJSON()
		return
	}else {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "登录失败！账号或者密码不正确",
			Info: "",
		}
		this.ServeJSON()
		return
	}
}


func (this *LoginController) CheckLoginStatus() {
	sess,err := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)//全局共享session
	if err != nil {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "session获取失败",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	userInfo := sess.Get(context.Background(),"LoginUserInfo").(models.UserInfo)
	if userInfo.Id == 0 {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "登录校验失败！",
			Info: "",
		}
		this.ServeJSON()
		return
	}
}

func (this *LoginController) Logout() {
	sess,err := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)//全局共享session
	if err != nil {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "session获取失败",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	sess.Delete(context.Background(),"LoginUserInfo")
	this.Data["json"] = models.RspComm{
		Code: 100,
		Msg:  "退出成功！！",
		Info: "",
	}
	this.ServeJSON()
}