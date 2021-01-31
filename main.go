package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

/ / Define the structure (xorm supports bidirectional mapping)

type User struct {
	Account string `json:"account" xorm:"not null pk default '' VARCHAR(32)"`
	Password  string `json:"passwd" xorm:"not null default '' VARCHAR(32)"`
	Email   string `json:"email" xorm:"not null default '' VARCHAR(32)"`
}
/ / Define the orm engine
var engine  *xorm.Engine



/ / Create an orm engine
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:newpass@dobol(127.0.0.1:3306)/tbuji?charset=utf8")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if  err := engine.Sync2(new(User)); err != nil {
		log.Fatal("Data table synchronization failed:", err)
	}

}


func main(){

	Install("123456","123456","790379511@qq.com")
	Find("123456")
	update("123456","123456","1234567")
}

func Install (account,password,email string){
	str1 := new(User)
	str1.Account = account
	str1.Password = password
	str1.Email = email
	install, err := engine.Insert(str1)
	if err != nil {
		log.Fatal("registration failed:", err)
	}
	fmt.Println(install)

}

func Find (account string) {
	a := new(User)
	str, err := engine.ID(account).Get(a)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	fmt.Println(str)
	fmt.Println(a.Account,a.Password,a.Email)

}

func update (account,oldpwd,newpwd string) {
	a := new(User)
	str, err := engine.ID(account).Get(a)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	if str == false {
		fmt.Println("Username does not exist!")
		return
	}

	if a.Password != oldpwd {
		fmt.Println("The old password is different!")
		return
	}
	a.Password = newpwd
	addr , err := engine.ID(account).Update(a)
	if err != nil{
		fmt.Println("fail to edit!")
	}

	fmt.Println("Successfully modified!",addr)
}