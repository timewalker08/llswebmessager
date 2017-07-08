package models

import (
    "time"
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id          int64
    Name    string
    PasswordMd5 string
    CreatedAt   time.Time
    Friend      []*Friend `orm:"reverse(many)"`
}

type Friendstatus struct {
    Id          int
    StatusName  string
}

type Friend struct {
    Id           int
    User         *User `orm:"rel(fk)"`
    Friend       *User `orm:"rel(fk)"`
    Friendstatus *Friendstatus `orm:"rel(fk)"`
    CreatedAt   time.Time
}

func init() {
    orm.RegisterModel(new(User), new(Friendstatus), new(Friend))
}