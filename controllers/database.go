package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

type DBController struct {
	beego.Controller
}

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3307)/tmp?charset=utf8", 30)
}

func (c *DBController) CreateTable() {
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *DBController) InsertData() {
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}

func (c *DBController) UpdateData() {
	o := orm.NewOrm()
	profile := &Profile{Id: 1, Age: 30}
	user := &User{Id: 1, Name: "Your", Profile: profile}

	fmt.Println(o.Update(user))
}

func (c *DBController) RawQuery() {
	o := orm.NewOrm()
	re, b := o.Raw("UPDATE user SET name = ? WHERE id = ?", "\"soo\"", "1").Exec()
	log.Println(re, b)

}

func (c *DBController) Delete() {
	o := orm.NewOrm()
	p := &Profile{Id: 1}
	o.Delete(p)
}
