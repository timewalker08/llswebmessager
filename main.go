package main

import (
    //"fmt"
    //"time"
    //"llswebmessager/models"
    //"llswebmessager/tools"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    
    _ "llswebmessager/routers"
    "github.com/astaxie/beego"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:liulishuo1@tcp(52.187.20.17:3306)/wmtest?charset=utf8")
}

func main() {
	beego.Run()
}

