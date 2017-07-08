package controllers

import (
    "fmt"
    "llswebmessager/models"
    "github.com/astaxie/beego"
)

type AccountController struct {
    beego.Controller
}

type formuser struct {
    Name  string `form:"username"`
    Password   string         `form:"password"`
    ConfirmPassword string    `form:"confirmpassword"`
}

// [Get]
func (this *AccountController) Register() {
    CheckError(this)
    this.TplName = "register.tpl"
}

// [Get]
func (this *AccountController) Login() {
    CheckError(this)
    this.TplName = "login.tpl"
}

func CheckError (c *AccountController) {
    var err string
    c.Ctx.Input.Bind(&err, "err")
    if err != "" {
      c.Data["HasError"] = true
      c.Data["ErrorMsg"] = err
    }
}

// [post] register user request
func (this *AccountController) RegisterUser() {
    u := formuser{}
    if err := this.ParseForm(&u); err != nil {
        this.Redirect("/account/register?err=" + err.Error(), 302)
    }
    saveFormData(this, &u)
    if err := checkRegisterUserInfo(&u); err != nil {
        this.Redirect("/account/register?err=" + err.Error(), 302)
    }
    
    // create user if not exist by name
    created, user, err := models.CreateUserIfNotExistByName(u.Name, u.Password)
    if (!created) {
        this.Redirect("/account/register?err=" + err.Error(), 302)
    } else {
	    this.SetSession(models.LoginSessionKey, &models.AccountManager{User: user})
	}
    
    this.Redirect("/friend/list", 302)
}

// [post] login request
func (this *AccountController) LoginUser() {
    u := formuser{}
    if err := this.ParseForm(&u); err != nil {
        this.Redirect("/account/login?err=" + err.Error(), 302)
    }
    saveFormData(this, &u)
    ret, user, err := models.CheckUserPassword(u.Name, u.Password);
	if !ret {
        this.Redirect("/account/login?err=" + err.Error(), 302)
    } else {
	    this.SetSession(models.LoginSessionKey, &models.AccountManager{User: user})
	}
    
    this.Redirect("/friend/list", 302)
}

func checkRegisterUserInfo (fuser *formuser) error {
    var errstr string = ""
    if (fuser.Name == "") {
	    errstr = "Name should not be empty. Please input a name."
    } else if (len(fuser.Password) < 6) {
	    errstr = "Password should be longer than 6 characters."
	} else if (fuser.Password != fuser.ConfirmPassword) {
	    errstr = "The password and confirmation password do not match."
	}
	if errstr != "" {
	    fmt.Println("invalid register user info: %s", errstr)
	    return &models.InvalidUserInfoWhenRegister{ErrorMessage: errstr}
	}
    return nil
}

func saveFormData (c *AccountController, u *formuser) {
    c.Data["Name"] = u.Name
    c.Data["Password"] = u.Password
    c.Data["ConfirmPassword"] = u.ConfirmPassword
}