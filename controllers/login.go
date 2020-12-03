package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	username, _ := c.Ctx.GetSecureCookie("123698741", "username")
	c.Data["username"] = username
	c.TplName = "login.tpl"
}

func (c *LoginController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)

	// 把 username 保存到 cookie 中
	c.Ctx.SetSecureCookie("123698741", "username", username)

	// 执行重定向跳转
	// c.Redirect("/", 302)

	c.Ctx.Redirect(302, "/cms_123456.html")
}

// 删除 cookie
func (c *LoginController) DoLoginOut() {
	c.Ctx.SetSecureCookie("123698741", "username", "", 0)

	// 执行重定向跳转
	c.Ctx.Redirect(302, "/cms_123456.html")
}
