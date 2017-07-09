package models

import (
    "fmt"
	"strconv"
    "time"
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
	if _, err := o.Update(msg, "Messagestatus"); err != nil {
	    fmt.Printf("Error when updating message status: %s\n", err.Error())
        return &CommonMessageError{ErrorMessage: fmt.Sprintf("Some error occurred when updating message. Error detail: %s", err.Error())}
    }
	return nil
}

func (this *MessageManager) GetLastReadTime(fromId int) (*time.Time, error) {
    if (this.User == nil) {
        return nil, &CommonError {ErrorMsg: "User is not set!"}
    }
	
    o := orm.NewOrm()
	
	lrmt := Lastreadmessagetime{From: &User{Id: fromId}, To: this.User}
	err := o.Read(&lrmt, "From", "To")
	if err != nil {
	    return nil, &CommonMessageError{ErrorMessage: fmt.Sprintf("Some error occurred when getting last read time. Error detail: %s", err.Error())}
	}
	return &lrmt.Lastreadtime, nil
}

func (this *MessageManager) SetLastReadTime(fromId int, lrt time.Time) error {
    if (this.User == nil) {
        return &CommonError {ErrorMsg: "User is not set!"}
    }
	
	lrmt := Lastreadmessagetime{From: &User{Id: fromId}, To: this.User, Lastreadtime: lrt}
	cplrmt := lrmt
	o := orm.NewOrm()
	err := o.Begin()
	created, _, err := o.ReadOrCreate(&lrmt, "From", "To");
	if err == nil && !created {
        cplrmt.Id = lrmt.Id
        _, err = o.Update(&cplrmt)
    }
	if err == nil {
        err = o.Commit()
    } else {
        err = o.Rollback()
    }
	
	if err != nil {
	    return &CommonMessageError{ErrorMessage: fmt.Sprintf("Some error occurred when setting last read time. Error detail: %s", err.Error())}
	}
	return nil
}

func (this *MessageManager) GetUnReadMessageCount() (*map[int]int, error) {
    if (this.User == nil) {
        return nil, &CommonError {ErrorMsg: "User is not set!"}
    }
	o := orm.NewOrm()
	res := make(orm.Params)
	queryFormat := `
	    SELECT m.from_id, count(1) as C
		FROM message m
		LEFT JOIN lastreadmessagetime lt
		ON m.from_id = lt.from_id AND m.to_id = lt.to_id
		WHERE m.to_id = %d
		AND m.created_at >= IFNULL(lt.lastreadtime, m.created_at)
		AND m.messagestatus_id = %d
		GROUP BY m.from_id
		`
    o.Raw(fmt.Sprintf(queryFormat, this.User.Id, MessagestatusNormal)).RowsToMap(&res, "from_id", "C")
	mmp := make(map[int]int)
	for key := range res {
	    ik, _ := strconv.Atoi(key)
		iv, _ := strconv.Atoi(res[key].(string))
		fmt.Printf("GetUnReadMessageCount: user_id: %d, unread count: %d\n", ik, iv)
	    mmp[ik] = iv
	}
	return &mmp, nil
}

// TODO: support pagination
func (this *MessageManager) GetUnReadMessagesByPage(fromId int, page int, pageSize int) ([]*Message, error) {
    if (this.User == nil) {
        return nil, &CommonError {ErrorMsg: "User is not set!"}
    }
	
	if page <= 0 || pageSize <= 0 {
	    return nil, &InvalidPaginationPara{Page: page, PageSize: pageSize}
	}
	
	lrmt, _ := this.GetLastReadTime(fromId)
	
	toId := this.User.Id
	
	cond := orm.NewCondition()
    cond1 := cond.And("from_id", fromId).And("to_id", toId)
	cond2 := cond.And("from_id", toId).And("to_id", fromId)
	cond3 := cond1.OrCond(cond2)
	
	o := orm.NewOrm()
	var ms []*Message
	qs := o.QueryTable("message")
	qs = qs.SetCond(cond3)
	
	if lrmt == nil {
	    qs.Filter("Messagestatus", NormalMessagestatus).OrderBy("CreatedAt").RelatedSel("From").All(&ms)  //
	} else {
        qs.Filter("Messagestatus", NormalMessagestatus).RelatedSel("From").OrderBy("CreatedAt").All(&ms)  //Filter("created_at__gte", &lrmt).
	}
	fmt.Printf("Get %d messages\n", len(ms))
	if lrmt != nil && len(ms) > 0 {
	    fmt.Printf("lrmt: %s, message create at: %s\n", (*lrmt).Format("2006-01-02 15:04:05"), ms[0].CreatedAt.Format("2006-01-02 15:04:05"))
	}
    return ms, nil
}