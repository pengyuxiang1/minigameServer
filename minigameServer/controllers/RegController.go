package controllers

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"io"
	"minigameServer/models"
	"regexp"
	"strconv"
	"time"
)

type RegController struct {
	web.Controller
}

//注册处理
func (this *RegController) Post() {
	username := this.Ctx.Request.Form.Get("username")
	accountId := this.Ctx.Request.Form.Get("account_id")
	password := this.Ctx.Request.Form.Get("password")
	usererr := checkUsername(username)
	if usererr == false {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "username error, Please to again",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	passerr := checkPassword(password)
	if passerr == false {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "Password error, Please to again",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	accountidErr := checkAccountId(accountId)
	if accountidErr == false {
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "account_id error, Please to again",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()//密文密码
	now := time.Now()

	userInfo,err := models.GetUserInfoByName(username)
	if err != nil {
		fmt.Println("GetUserInfoByName执行出错")
		this.Data["json"] = models.RspComm{
			Code: 200,
			Msg:  "GetUserInfoByName执行出错！",
			Info: "",
		}
		this.ServeJSON()
		return
	}
	if userInfo.UserName == "" {
		users:= &models.UserInfo{
			UserName:  username,
			AccountId: accountId,
			Password:  newPass,
			Created:   now,
			Updated:   now,
		}
		userInfoRsp,err := models.AddUserInfo(users)
		if err != nil {
			fmt.Println("存储新账户失败")
			this.Data["json"] = models.RspComm{
				Code: 200,
				Msg:  "注册失败，新用户写入异常！",
				Info: "",
			}
			this.ServeJSON()
			return
		}
		//登录成功设置session
		//seesionId ,err:= getSessionId(userInfoRsp)
		//if err != nil {
		//	fmt.Println("获取sessionid失败")
		//	return
		//}
		//this.StartSession()
		//this.SetSession(seesionId, userInfoRsp)
		sess,err := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)//全局共享session
		sess.Set(context.Background(),"LoginUserInfo",userInfoRsp)
		this.Data["json"] = models.RspComm{
			Code: 100,
			Msg:  "注册成功，已自动登录！",
			Info: "",
		}
		this.ServeJSON()
		return
	} else {
		this.Data["UsernameErr"] = "用户名已经存在"
	}
}

func getSessionId(userInfo *models.UserInfo) (string,error) {
	sessionIdGene := md5.New()
	_ ,err:=io.WriteString(sessionIdGene, strconv.Itoa(int(userInfo.Id))+time.Now().String())
	if err != nil {
		return "",err
	}
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", sessionIdGene.Sum(nil))
	sessionId := buffer.String()//密文密码
	return sessionId,nil
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkAccountId(accountid string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", accountid); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}