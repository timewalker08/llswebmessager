package controllers

import (
    "fmt"
    "llswebmessager/models"
    "github.com/astaxie/beego"
    //"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type FriendController struct {
    beego.Controller
}

func (this *FriendController) List() {
    am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
    
    fs, err := am.FriendManager.GetFriends()
    if err != nil {
	    fmt.Printf("Error: %s\n", err.Error())
	} else {
	    fmt.Printf("get %d friends\n", len(fs))
	}
	mmp, _ := am.MessageManager.GetUnReadMessageCount()
	
	var fcs []*models.FriendWithUnReadCount// = make([]*models.FriendWithUnReadCount)
	if (fs != nil && len(fs) > 0) {
	    for _, f := range fs {
		    count := 0
			if cc, ok := (*mmp)[f.Friend.Id]; ok {
			    count = cc
			}
			fmt.Printf("Friend id: %d,Friend name: %s, unreadcount: %d\n", f.Friend.Id, f.Friend.Name, count)
	        fcs = append(fcs, &models.FriendWithUnReadCount{Friend: f, UnreadCount: count})
		}
	}

	this.Data["UserName"] = am.User.Name
    this.Data["FriendList"] = fcs;
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
    am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
	if am == nil {
	    return;
	}
    err := am.AddNewFriend(name)
    if err == nil {
      this.Data["json"] = models.WebApiResult{Code: 0}
    }else {
      this.Data["json"] = models.WebApiResult{Code: -1, Msg: err.Error()}
    }
    this.ServeJSON()
}

// [Post] Web api for deleting friend by name.
// Url: friend/delete
func (this *FriendController) DeleteFriend() {
    var name string
    this.Ctx.Input.Bind(&name, "name")
    am := this.GetLoginAMAndRedictToLoginPageIfNotLoggedin()
    am.DeleteFriendByName(name)
    this.Data["json"] = models.WebApiResult{Code: 0}
    this.ServeJSON()
}

func (this *FriendController)GetLoginAM() *models.AccountManager {
    am := this.GetSession(models.LoginSessionKey)
    if (am == nil) {
        return nil
    }
    return am.(*models.AccountManager)
}

func (this *FriendController)GetLoginAMAndRedictToLoginPageIfNotLoggedin() *models.AccountManager {
    var am = this.GetLoginAM()
    if am == nil {
        this.Redirect("/account/login?err=Please login first.", 302)
        return nil
    }
    return am
}