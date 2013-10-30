package main

import (
	"admin/controllers"
	"admin/controllers/rbac"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	//orm.Debug = true
	fmt.Println("Starting....")
	orm.RegisterDataBase("default", "mysql", "root:root@/admin?charset=utf8")
	//CreateDB()

	beego.Router("/", &controllers.MainController{})
	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user", &rbac.UserController{}, "*:Index")

	beego.Run()
}

func CreateDB() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
