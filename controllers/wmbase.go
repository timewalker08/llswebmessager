package controllers

import (
    "fmt"
	"llswebmessager/models"
	"github.com/astaxie/beego"
)

type WMBaseController struct {
	beego.Controller
}

func (this *WMBaseController) IsLogined() bool {
    fmt.Println("------ in is logined")
	am := this.GetSession(models.LoginSessionKey)
	return am != nil
}