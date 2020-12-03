package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	// 设置cookie， 第3个参数，是时间(过期时间)，单位是秒
	// SetCookie 不可设置中文
	// c.Ctx.SetCookie("cookieUsername", "Tom", 20)

	// 设置cookie，第4个参数，访问的路径
	// c.Ctx.SetCookie("cookieUsername", "Tom", 20, "/view")

	// 设置cookie， 第5个参数，cookie的路径 Domain(二级域名下也可以访问cookie)
	// itying.com  www.itying.com   a.itying.com   b.itying.com
	// c.Ctx.SetCookie("cookieUsername", "Tom", 20, "/", ".itying.com")

	// 设置cookie， 第6个参数，secure，当 secure 值为 true 时，cookie 在 HTTP 中无效， 在 HTTPS 中才有效
	// secure 默认为 false

	// 设置cookie， 第7个参数， httpOnly，微软对cookie的扩展，防止 XSS 攻击（通过JS脚本等程序读取cookie）
	// httpOnly 默认 false

	// SetSecureCookie
	// 第一个参数为秘钥
	c.Ctx.SetSecureCookie("123698741", "cookieUsername", "张三", 20)

	c.SetSession("username", "王五")

	c.TplName = "index.tpl"
}
