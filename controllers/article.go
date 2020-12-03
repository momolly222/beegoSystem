package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Data["content"] = "baby"
	c.TplName = "article.html"
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻")
}

func (c *ArticleController) EditArticle() {
	// 获取get传值

	// GetString
	// id := c.GetString("id")
	// beego.Info(id)

	// GetInt
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("传入参数错误")
		return
	}
	fmt.Printf("类型：%T", id)
	c.Ctx.WriteString("编辑新闻--")
}
