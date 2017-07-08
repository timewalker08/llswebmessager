package routers

import (
    "webmessagertest/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.AccountController{})
    beego.AutoRouter(&controllers.FriendController{})
    beego.Router("/account/registeruser",&controllers.AccountController{},"post:RegisterUser")
	beego.Router("/account/loginuser",&controllers.AccountController{},"post:LoginUser")
	beego.Router("/friend/queryname",&controllers.FriendController{},"get:QueryName")
	beego.Router("/friend/add",&controllers.FriendController{},"post:AddFriend")
	beego.Router("/friend/delete",&controllers.FriendController{},"post:DeleteFriend")
}
