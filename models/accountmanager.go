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
    FriendManager *FriendManager
    MessageManager *MessageManager
}

func (this *AccountManager) AddNewFriend(name string) error {
    var friend Friend
    friend.User = this.User
    friend.Friend = QueryUserByName(name)
    if (friend.Friend == nil) {
        return &AccountNotExistError{AccountName: name}
    }
    friend.CreatedAt = time.Now()
    friend.Friendstatus = NormalFriendStatus
    if _, err := this.FriendManager.CreateOrUpdateFriend(&friend); err != nil {
        return &AddFriendFailedError{UserName: this.User.Name, FriendName: name}
    }
    return nil
}

func (this *AccountManager) DeleteFriendByName(name string) error {
    friend := QueryUserByName(name)
    if (friend == nil) {
        return &AccountNotExistError{AccountName: name}
    }
    
    return this.FriendManager.DeleteFriend(this.User, friend)
}

func (this *AccountManager) SendMessage(name string, msgstr string) (*Message, error) {
    friend := QueryUserByName(name)
    if friend == nil {
        return nil, &AccountNotExistError{AccountName: name}
    }
    
    fd := Friend{User: friend, Friend: this.User, Friendstatus: NormalFriendStatus, CreatedAt: time.Now()}
    if _, err := this.FriendManager.CreateOrUpdateFriend(&fd); err != nil {
        return nil, &AddFriendFailedError{UserName: name, FriendName: this.User.Name}
    }
    
    msg := &Message{From: this.User, To: friend, Msg: msgstr, Messagestatus: NormalMessagestatus, CreatedAt:time.Now()}
    return msg, this.MessageManager.SendMessage(msg)
}

func (this *AccountManager) DeleteMessage(msgId int) error {
    msg, err := this.MessageManager.GetMessageById(msgId)
    if err == nil {
        fmt.Printf("From id: %d, user id: %d\n", msg.From.Id, this.User.Id)
        if msg.From.Id != this.User.Id {
            return &CommonMessageError{ErrorMessage: fmt.Sprintf("User can only delete messages sent by oneself")}
        }
        msg.Messagestatus = DeletedMessagestatus
        return this.MessageManager.UpdateMessageStatus(msg)
    } else {
        return err
    }
}

func (this *AccountManager) GetMessagesByPage(fromName string, page int, pageSize int) ([]*Message, error) {
    friend := QueryUserByName(fromName)
    if friend == nil {
        return nil, &AccountNotExistError{AccountName: fromName}
    }
    
    msgs, err := this.MessageManager.GetMessagesByPage(friend.Id, page, pageSize)
    if err == nil {
        err = this.MessageManager.SetLastReadTime(friend.Id, time.Now())
    }
    if err != nil {
        return nil, &CommonMessageError{ErrorMessage: fmt.Sprintf("Some error occurred when getting unread messages, error detail: %s", err.Error())}
    }
    return msgs, nil
}

func CreateUserIfNotExistByName (name string, password string) (bool, *User, error) {
    o := orm.NewOrm()
    user := &User{Name: name, PasswordMd5: password, CreatedAt: time.Now()}  //TODO: calculate MD5
    created, _, err := o.ReadOrCreate(user, "Name")
    if !created && err == nil {
        return false, nil, &NameAlreadyUsedWhenRegisterError{UserName: name}
    }

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