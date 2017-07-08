package controllers

import (
    //"fmt"
	"llswebmessager/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    am := this.GetSession(models.LoginSessionKey)
	if am != nil {
	    this.Redirect("/friend/list", 302)
	}
	this.TplName = "index.tpl"
}
