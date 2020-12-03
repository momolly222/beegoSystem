package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {
	c.Data["myunix"] = 1587880013
	cookieUsername, _ := c.Ctx.GetSecureCookie("123698741", "cookieUsername")
	c.Data["cookieUser"] = cookieUsername
	c.TplName = "user.html"
}

// 获取参数
// 方法一：
// func (c *UserController) DoUser() {
// 	name := c.GetString("username")
// 	password := c.GetString("password")
// 	c.Ctx.WriteString("用户---" + name + password)
// }

// 方法二：
type User struct {
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hhh"`
}

func (c *UserController) DoUser() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		beego.Info(err)
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析post数据成功")
}

// 在beego中，如果我们需要返回一个json数据的话，需要把数据放在结构体中
func (c *UserController) GetUser() {
	u := User{
		Username: "张三",
		Password: "123456",
		Hobby:    []string{"1", "3"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}
