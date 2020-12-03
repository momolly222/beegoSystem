package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) DoRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	// 执行重定向跳转
	// c.Redirect("/", 302)

	c.TplName = "success.html"
}
