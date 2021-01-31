package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"log"
)

// Define the structure (xorm supports bidirectional mapping)

type User struct {
	Name      string `json:"name"   xorm:"not null pk default '' VARCHAR(82)"`
	Password  string `json:"passwd" xorm:"not null default '' VARCHAR(32)"`
	Email     string `json:"email"  xorm:"not null default '' VARCHAR(32)"`
}
// Define the orm engine
var engine *xorm.Engine



// Create an orm engine
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:newpass@tcp(127.0.0.1:3306)/dobol?charset=utf8")
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

func Install(name,passwd,email string){
	str1 := new(User)
	str1.Name = name
	str1.Password = passwd
	str1.Email = email
	install, err := engine.Insert(str1)
	if err != nil {
		log.Fatal("registration failed:", err)
	}
	fmt.Println(install)

}

func Find(name string) {
	a := new(User)
	str, err := engine.ID(name).Get(a)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	fmt.Println(str)
	fmt.Println(a.Name,a.Password,a.Email)

}

func update(name,oldpwd,newpwd string) {
	a := new(User)
	str, err := engine.ID(name).Get(a)
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
	addr , err := engine.ID(name).Update(a)
	if err != nil{
		fmt.Println("fail to edit!")
	}

	fmt.Println("Successfully modified!",addr)
}
