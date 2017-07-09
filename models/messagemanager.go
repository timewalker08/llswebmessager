package models

import (
    "fmt"
    //"time"
    //"llswebmessager/tools"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type MessageManager struct {
    User *User
}

func (this *MessageManager) SendMessage(msg *Message) error {
	o := orm.NewOrm()
	_, err := o.Insert(msg)
	fmt.Printf("New message id: %d", msg.Id)
	if err != nil {
	    return &SendMessageError{ErrorDetail:err.Error()}
	}
	return nil
}

func (this *MessageManager) UpdateMessageStatus(msg *Message) error {
    o := orm.NewOrm()
	if _, err := o.Update(&msg, "Messagestatus"); err != nil {
        return &CommonMessageError{ErrorMessage: fmt.Sprintf("Some error occurred when updating message. Error detail: %s", err.Error())}
    }
	return nil
}

func (this *MessageManager) GetUnReadMessagesByPage(fromId int, page int, pageSize int) ([]*Message, error) {
    if (this.User == nil) {
        return nil, &CommonError {ErrorMsg: "User is not set!"}
    }
	
	if page <= 0 || pageSize <= 0 {
	    return nil, &InvalidPaginationPara{Page: page, PageSize: pageSize}
	}
	
	//toId := this.User.Id
	
	//cond := orm.NewCondition()
    //cond1 := cond.And("from_id", fromId).And("to_id", toId)
	//cond2 := cond.And("from_id", toId).And("to_id", fromId)
	//cond3 := cond1.OrCond(cond2)
	
	o := orm.NewOrm()
	qs := o.QueryTable("message")
	//qs = qs.SetCond(cond1)
	
    var ms []*Message
    qs.RelatedSel().All(&ms)
	fmt.Printf("Get %d messages\n", len(ms))
    return ms, nil
}