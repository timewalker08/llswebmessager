package models

import (
    "fmt"
    "time"
    //"llswebmessager/tools"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type AccountManager struct {
    User *User
	MessageManager *MessageManager
}

func (this *AccountManager) AddNewFriend(name string) error {
    if (this.User == nil) {
        return &CommonError {ErrorMsg: "User is not set!"}
    }
    
    var friend Friend
    friend.User = this.User
    friend.Friend = QueryUserByName(name)
	if (friend.Friend == nil) {
	    return &AccountNotExistError{AccountName: name}
	}
    friend.CreatedAt = time.Now()
    friend.Friendstatus = NormalFriendStatus
    if !createOrUpdateFriend(&friend) {
	    return &AddFriendFailedError{UserName: this.User.Name, FriendName: name}
	}
    return nil
}

func (this *AccountManager) DeleteFriendByName(name string) (bool, error) {
    if (this.User == nil) {
        ce := CommonError {ErrorMsg: "User is not set!"}
        return false, &ce
    }
    
    var friend Friend
    friend.User = this.User
    friend.Friend = QueryUserByName(name)
	if (friend.Friend == nil) {
	    return false, &AccountNotExistError{AccountName: name}
	}
    
    o := orm.NewOrm()
    o.Read(&friend, "User", "Friend")
    friend.Friendstatus = DeletedFriendStatus
    fmt.Printf("user id: %d, friend id: %d, status id: %d \n", friend.User.Id, friend.Friend.Id, friend.Friendstatus.Id)
    o.Update(&friend)
    fmt.Printf("friend %s of %s is deleted\n", name, this.User.Name)
    
    return true, nil
}

func (this *AccountManager) SendMessage(name string, msgstr string) (*Message, error) {
    friend := QueryUserByName(name)
	if friend == nil {
	    return nil, &AccountNotExistError{AccountName: name}
	}
	
	msg := new(Message)
	msg.From = this.User
	msg.To = friend
	msg.Msg = msgstr
	msg.Messagestatus = NormalMessagestatus
	msg.CreatedAt = time.Now()
	
	return msg, this.MessageManager.SendMessage(msg)
}

func (this *AccountManager) DeleteMessage(msgId int) error {
    msg := new(Message)
	msg.Id = msgId
	msg.Messagestatus = DeletedMessagestatus
    return this.MessageManager.UpdateMessageStatus(msg)
}

func (this *AccountManager) GetUnReadMessagesByPage(fromName string, page int, pageSize int) ([]*Message, error) {
    friend := QueryUserByName(fromName)
	if friend == nil {
	    return nil, &AccountNotExistError{AccountName: fromName}
	}
	
    return this.MessageManager.GetUnReadMessagesByPage(int(friend.Id), page, pageSize)
}

func createOrUpdateFriend(friend *Friend) bool {
    o := orm.NewOrm()
    copyf := *friend
    err := o.Begin()
    
    created, _, err := o.ReadOrCreate(friend, "User", "Friend");
    if err == nil && !created {
        copyf.Id = friend.Id
        _, err = o.Update(&copyf)
    }
    
    if err == nil {
        err = o.Commit()
    } else {
        err = o.Rollback()
    }
    fmt.Println(err)
    
    return created
}

func CreateUserIfNotExistByName (name string, password string) (bool, *User, error) {
    o := orm.NewOrm()
    user := new(User)
	user.Name = name
    user.PasswordMd5 = password                                  //TODO: calculate MD5
    user.CreatedAt = time.Now()
    created, _, err := o.ReadOrCreate(user, "Name")
	if !created {
	    return false, nil, &NameAlreadyUsedWhenRegisterError{UserName: name}
	}

    fmt.Println(err, user.Id)
	
    return created, user, err
}

func CheckUserPassword (name string, password string) (bool, *User, error) {
    o := orm.NewOrm()
    user := new(User)
	user.Name = name
    err := o.Read(user, "Name")
    if (err == nil && user.PasswordMd5 == password) {
        return true, user, nil
    } else {
	    return false, nil, &LoginError{UserName: name, Password: password}
    }
}

func QueryUserByName (name string) *User {
    o := orm.NewOrm()
    user := new(User)
    user.Name = name
    err := o.Read(user, "Name")
    if (err == nil) {
        return user
    }
    return nil
}