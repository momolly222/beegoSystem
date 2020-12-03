package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

// 路由的伪静态
func (c *CmsController) Get() {
	cmsid := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms详情---" + cmsid)
}
