package models

import (
    "fmt"
    //"time"
    //"llswebmessager/tools"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type FriendManager struct {
    User *User
}

func (this *FriendManager) GetFriends() ([]*Friend, error) {
	o := orm.NewOrm()
	var fs []*Friend
	_, err := o.QueryTable("friend").Filter("User", this.User.Id).Filter("Friendstatus", NormalFriendStatus).RelatedSel(1).All(&fs)
	if err != nil {
	    fmt.Printf("UserId: %d, Error detail: %s\n", this.User.Id, err.Error())
	    return nil, &CommonError{ErrorMsg: fmt.Sprintf("An error occurred when getting friend. Error detail: %s", err.Error())}
	}
	return fs, nil
}