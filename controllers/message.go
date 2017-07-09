package controllers

import (
    "fmt"
    "llswebmessager/models"
    "github.com/astaxie/beego"
    //"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type MessageController struct {
    beego.Controller
}

// [Get] Web api. used to get unread messages
// Url: /message/unread
func (this *MessageController) GetUnreadMessages() {
    var name string
    this.Ctx.Input.Bind(&name, "name")
    am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
	if am == nil {
	    return;
	}
    msgs, _ := am.GetUnReadMessagesByPage(name, 1, 100)
    this.Data["json"] = msgs
	if msgs != nil {
	    fmt.Printf("Get %d messages.\n", len(msgs))
	}
    this.ServeJSON()
}

// [Post] Web api for sending message to friend.
// Url: /message/new
// TODO: should put data in http body.
func (this *MessageController) SendMessage() {
    fmt.Println("--------in send message")
    var name string
	var msgstr string
	this.Ctx.Input.Bind(&name, "name")
	this.Ctx.Input.Bind(&msgstr, "msg")
    am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
	if am == nil {
	    return
	}
	_, err := am.SendMessage(name, msgstr)
	if err == nil {
	    fmt.Println("no error")
	    this.Data["json"] = models.WebApiResult{Code: 0}
	} else {
	    fmt.Printf("Some error: %s\n", err.Error())
	    this.Data["json"] = models.WebApiResult{Code: -1, Msg: err.Error()}
	}
    this.ServeJSON()
}

// [Post] Web api for removing message
// Url: /message/remove
func (this *MessageController) DeleteMessage() {
    fmt.Println("--------in delete message")
    var msgId int
	this.Ctx.Input.Bind(&msgId, "msgId")
	am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
	if am == nil {
	    return
	}
	err := am.DeleteMessage(msgId)
	if err == nil {
	    fmt.Println("no error")
	    this.Data["json"] = models.WebApiResult{Code: 0}
	} else {
	    fmt.Printf("Some error: %s\n", err.Error())
	    this.Data["json"] = models.WebApiResult{Code: -1, Msg: err.Error()}
	}
	this.ServeJSON()
}

//TODO: refactor
func (this *MessageController)GetLoginAM() *models.AccountManager {
    am := this.GetSession(models.LoginSessionKey)
    if (am == nil) {
        return nil
    }
    return am.(*models.AccountManager)
}

func (this *MessageController)GetLoginAMAndRedictToLoginPageIfNotLoggedin() *models.AccountManager {
    var am = this.GetLoginAM()
    if am == nil {
        this.Redirect("/account/login?err=Please login first.", 302)
        return nil
    }
    return am
}
