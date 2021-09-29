package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type UserInfo struct {
	Id   int64
	UserName string  `orm:"description(用户名)"`
	AccountId  string  `orm:"description(账号id)"`
	Password string  `orm:"description(密码)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	dbconn,err := web.AppConfig.String("DBConn")
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	db.Ping()
	Db = db
}

func Close() {
	if Db != nil {
		Db.Close()
	}
}

func AddUserInfo(userInfo *UserInfo) (*UserInfo, error) {
	var isql = "INSERT `user_info` SET user_name=?,account_id=?,password=?,created=?,updated=?"
	response := &UserInfo{
		Id:       0,
		UserName: "",
		AccountId:  "",
		Password: "",
		Created:  time.Time{},
		Updated:  time.Time{},
	}
	if Db == nil {
		return response, errors.New("connect mysql failed")
	}
	stmt, _ := Db.Prepare(isql)
	defer stmt.Close()
	res, err := stmt.Exec(userInfo.UserName, userInfo.AccountId, userInfo.Password, userInfo.Created, userInfo.Updated)
	if err == nil {
		response.Id, _ = res.LastInsertId()
		return response, nil
	}
	return response, nil
}


func GetUserInfoById(AccountID string) (*UserInfo, error) {
	var qsql = "SELECT * FROM user_info WHERE  account_id=?"
	response :=&UserInfo{}
	if AccountID != "" {
		if Db == nil {
			return response, errors.New("connect mysql failed")
		}
		stmt, err := Db.Prepare(qsql)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		rows, err := stmt.Query(AccountID)
		defer rows.Close()
		if err != nil {
			return response, err
		}
		for rows.Next() {
			err = rows.Scan(&response.Id, &response.UserName, &response.AccountId, &response.Password,
				&response.Created, &response.Updated)
			if err != nil {
				return nil, err
			}
		}
		return response, nil
	}
	return response, errors.New("Requset is non porinter")
}

func GetUserInfoByName(userName string) (*UserInfo, error) {
	var qsql = "SELECT * FROM user_info WHERE  user_name=?"
	response :=&UserInfo{}
	if userName != "" {
		if Db == nil {
			return response, errors.New("connect mysql failed")
		}
		stmt, _ := Db.Prepare(qsql)
		rows, err := stmt.Query(userName)
		defer rows.Close()
		if err != nil {
			return response, err
		}
		for rows.Next() {
			err = rows.Scan(&response.Id, &response.UserName, &response.AccountId, &response.Password,
				&response.Created, &response.Updated)
			if err != nil {
				return nil, err
			}
		}
		return response, nil
	}
	return nil, errors.New("Requset is non porinter")
}
