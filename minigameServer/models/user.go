package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User2 struct {
	Id   int
	UserName string  `orm:"description(用户名)"`
	Account  string  `orm:"description(账号id)"`
	Password string  `orm:"description(密码)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init(){
	orm.RegisterModel(new(User2))
}
