package main

import (
    //"fmt"
    //"time"
    //"llswebmessager/models"
    //"llswebmessager/tools"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    
    _ "llswebmessager/routers"
    "github.com/astaxie/beego"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:liulishuo1@tcp(52.187.20.17:3306)/wmtest?charset=utf8")
}

func main() {
    //user1 := &models.User{Id: 2}
	//user2 := &models.User{Id: 5}
    //o := orm.NewOrm()
    //o.Read(user1)
	//o.Read(user2)
    //am1 := &models.AccountManager{User: user1, MessageManager: &models.MessageManager{User: user1}}
	//am2 := &models.AccountManager{User: user2, MessageManager: &models.MessageManager{User: user2}}
    //for a := 0; a < 10; a++ {
	//    msg1 := fmt.Sprintf("%d Hello Bob, this is an auto reply message. %d", a, a)
    //    am1.SendMessage(user2.Name, msg1)
	//	fmt.Printf("%s send message to %s\n", user1.Name, user2.Name)
	//	time.Sleep(1000000000)
	//	
	//	msg2 := fmt.Sprintf("%d Hello Antony, this is an auto reply message. %d", a, a)
    //    am2.SendMessage(user1.Name, msg2)
	//	fmt.Printf("%s send message to %s\n", user2.Name, user1.Name)
    //    time.Sleep(1000000000)
    //}
	//
	//msgs, err := am1.GetUnReadMessagesByPage(user2.Name, 1, 100)
	//if err != nil {
	//    fmt.Println(err.Error())
	//} else {
	//    fmt.Printf("%d messages got\n", len(msgs))
	//    for _, msg := range msgs {
    //        fmt.Printf("Id: %d, FromId: %d, ToId: %d, Message: %s \n", msg.Id, msg.From.Id, msg.To.Id, msg.Msg)
    //    }
	//}
	//
	//cnt, err := o.QueryTable("message").GroupBy("from_id").Count() // SELECT COUNT(*) FROM USER
    //fmt.Printf("Count Num: %d, %s\n", cnt, err)
	
    beego.Run()
}

