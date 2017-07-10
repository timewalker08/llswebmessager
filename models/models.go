package models

import (
    "time"
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id          int
    Name        string
    PasswordMd5 string
    CreatedAt   time.Time
    Friend      []*Friend `orm:"reverse(many)"`
    Message     []*Message `orm:"reverse(many)"`
}

type Friendstatus struct {
    Id          int
    StatusName        string
}

type Friend struct {
    Id           int
    User         *User `orm:"rel(fk)"`
    Friend       *User `orm:"rel(fk)"`
    Friendstatus *Friendstatus `orm:"rel(fk)"`
    CreatedAt   time.Time
}

type FriendWithUnReadCount struct {
    Friend *Friend
    UnreadCount int
}

type Messagestatus struct {
    Id          int
    Name        string
}

type Message struct {
    Id          int
    From        *User `orm:"rel(fk)"`
    To          *User `orm:"rel(fk)"`
    Msg         string
    Messagestatus *Messagestatus `orm:"rel(fk)"`
    CreatedAt   time.Time
}

type Lastreadmessagetime struct {
    Id          int
    From        *User `orm:"rel(fk)"`
    To          *User `orm:"rel(fk)"`
    Lastreadtime time.Time
}

func init() {
    orm.RegisterModel(new(User), new(Friendstatus), new(Friend), new(Messagestatus), new(Message), new(Lastreadmessagetime))
}