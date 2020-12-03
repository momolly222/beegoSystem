package controllers

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["content"] = "hello，beego" // 绑定数据
	c.TplName = "goods.tpl"
}

func (c *GoodsController) DoAdd() { // post
	c.Ctx.WriteString("执行增加操作")
}

type Product struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) DoEdit() {
	p := Product{}
	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("获取数据失败")
		return
	}
	fmt.Printf("%#v", p)
	c.Ctx.WriteString("执行修改操作")
}

func (c *GoodsController) DoDelete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("参数错误")
	}
	c.Ctx.WriteString("执行删除操作" + strconv.Itoa(id))
}

// 接收 Post 传过来的 XML 数据
// 注意：一定要在配置文件中设置 copyrequestbody=true
/*
<?xml version="1.0" encoding="UTF-8"?>
<article>
<content type="string">我是张三</content>
<title type="string">张三</title>
</article>
*/
func (c *GoodsController) Xml() {
	p := Product{}

	// str := string(c.Ctx.Input.RequestBody)
	// beego.Info(str)

	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); e != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		fmt.Printf("%#v", p)
		c.Data["json"] = p
		c.ServeJSON()
	}
}
