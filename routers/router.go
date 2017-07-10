package routers

import (
    "llswebmessager/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.AccountController{})
    beego.AutoRouter(&controllers.FriendController{})
    beego.AutoRouter(&controllers.MessageController{})
    beego.Router("/account/registeruser",&controllers.AccountController{},"post:RegisterUser")
    beego.Router("/account/loginuser",&controllers.AccountController{},"post:LoginUser")
    beego.Router("/friend/queryname",&controllers.FriendController{},"get:QueryName")
    beego.Router("/friend/add",&controllers.FriendController{},"post:AddFriend")
    beego.Router("/friend/remove",&controllers.FriendController{},"post:DeleteFriend")
    beego.Router("/message/all",&controllers.MessageController{},"get:GetMessages")
    beego.Router("/message/new",&controllers.MessageController{},"post:SendMessage")
    beego.Router("/message/remove",&controllers.MessageController{},"post:DeleteMessage")
}
