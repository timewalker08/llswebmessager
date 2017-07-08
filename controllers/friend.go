package controllers

import (
	"llswebmessager/models"
    "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type FriendController struct {
    beego.Controller
}

func (this *FriendController) List() {
    am := GetUser(this)
	if am == nil {
	    this.Redirect("/account/login?err=Please login first.", 302)
	}
    
	var friendStatus models.Friendstatus
    friendStatus.Id = 1
	
    //query friends from user id
	o := orm.NewOrm()
    var fs []*models.Friend
    o.QueryTable("friend").Filter("User", am.User.Id).Filter("Friendstatus", &friendStatus).RelatedSel().All(&fs)

	this.Data["FriendList"] = fs;
    this.TplName = "friendlist.tpl"
}

// [Get] Web api for getting user by name. Used when adding friend
// Url: friend/queryname?name=xxx
func (this *FriendController) QueryName() {
    var name string
    this.Ctx.Input.Bind(&name, "name")
    user := models.QueryUserByName(name)
    this.Data["json"] = user
    this.ServeJSON()
}

// [Post] Web api for adding new friend by name. Used when adding friend
// Url: friend/add
func (this *FriendController) AddFriend() {
    var name string
    this.Ctx.Input.Bind(&name, "name")
    am := GetUser(this)
    err := am.AddNewFriend(name)
    if err == nil {
      this.Data["json"] = models.WebApiResult{Code: 0}
    }else {
      this.Data["json"] = models.WebApiResult{Code: 0, Msg: err.Error()}
    }
    this.ServeJSON()
}

// [Post] Web api for deleting friend by name.
// Url: friend/delete
func (this *FriendController) DeleteFriend() {
    var name string
    this.Ctx.Input.Bind(&name, "name")
	am := GetUser(this)
    am.DeleteFriendByName(name)
    this.Data["json"] = models.WebApiResult{Code: 0}
    this.ServeJSON()
}

func GetUser(ctrl *FriendController) *models.AccountManager {
    am := ctrl.GetSession(models.LoginSessionKey)
	if (am == nil) {
	    return nil
	}
	return am.(*models.AccountManager)
}

