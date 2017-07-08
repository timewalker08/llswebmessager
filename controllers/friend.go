package controllers

import (
    "fmt"
	"llswebmessager/models"
    "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type FriendController struct {
    beego.Controller
}

func (this *FriendController) List() {
    fmt.Println("In friend list")
    
	var friendStatus models.Friendstatus
    friendStatus.Id = 1
	
    //query friends from user id
	o := orm.NewOrm()
    var fs []*models.Friend
    num, err := o.QueryTable("friend").Filter("User", 2).Filter("Friendstatus", &friendStatus).RelatedSel().All(&fs)
    if err == nil {
        fmt.Printf("%d friends read\n", num)
        for _, ff := range fs {
            fmt.Printf("Id: %d, UserName: %s, FriendName: %s \n", ff.Id, ff.User.Name, ff.Friend.Name)
        }
    }
	
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
    var user = models.User{Id: 2}
    var am = models.AccountManager{User: &user}
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
    fmt.Println("--------")
    var name string
    this.Ctx.Input.Bind(&name, "name")
    var user = models.User{Id: 2}
    var am = models.AccountManager{User: &user}
    am.DeleteFriendByName(name)
    this.Data["json"] = models.WebApiResult{Code: 0}
    this.ServeJSON()
}