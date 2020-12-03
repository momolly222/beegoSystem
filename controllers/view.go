package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

type Article struct {
	Title   string
	Content string
}

func (c *ViewController) Get() {
	// 1、模板中绑定基本数据 字符串 数值 布尔值
	c.Data["website"] = "beego教程"
	c.Data["num"] = 12
	c.Data["falg"] = true

	// 2、模板中绑定结构体数据
	a := Article{
		Title:   "我是beego教程",
		Content: "beego仿小米商城项目",
	}
	c.Data["article"] = a

	// 3、模板中循环遍历 range
	c.Data["articleList"] = []string{"php", "golang", "python"}

	// 4、模板中循环遍历 Map
	userinfo := make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 28
	userinfo["sex"] = "男"
	c.Data["userinfo"] = userinfo

	// 5、模板中循环遍历结构体类型切片
	articleGroup := []Article{
		{
			Title:   "新闻1",
			Content: "新闻1内容",
		},
		{
			Title:   "新闻2",
			Content: "新闻2内容",
		},
		{
			Title:   "新闻3",
			Content: "新闻3内容",
		},
	}
	c.Data["articleGroup"] = articleGroup

	// 6、模板中条件判断
	c.Data["isLogin"] = true
	c.Data["isHome"] = false

	// 9、格式化时间
	now := time.Now()
	c.Data["now"] = now

	// 10、字符串的截取
	str := "这是一个字符串的截取"
	c.Data["str"] = str

	// 11、字符串与html互换渲染
	html := "<h1>这是一个h1标题</h1>"
	c.Data["html"] = html

	// 13、map_get 获取 map 值
	m := make(map[string]interface{})
	m["a"] = 1
	m["b"] = map[string]float64{
		"c": 4,
		"d": 5,
	}
	c.Data["m"] = m

	// 15、把时间戳转换为格式化字符串
	c.Data["mytime"] = 1587880013

	c.Data["myunix"] = "2020-09-22 00:06:34"
	c.Data["mytime2"] = 1600704394

	// 16、获取配置文件里面的信息
	// beego.AppConfig.String 获取一个值
	appname := beego.AppConfig.String("appname")
	beego.Info(appname)
	// beego.AppConfig.Strings 获取多个值，返回一个切片
	admin := beego.AppConfig.String("admin")
	admin2 := beego.AppConfig.Strings("admin")
	beego.Info(admin)
	beego.Info(admin2)

	configUser := beego.AppConfig.String("user")
	beego.Info(configUser)

	// 17、获取default.go中的cookie信息
	cookieUsername, _ := c.Ctx.GetSecureCookie("123698741", "cookieUsername")
	c.Data["cookieUser"] = cookieUsername

	// 18、获取 session 信息
	// 方法一：
	sessionUser := c.GetSession("username")
	// 方法二：
	// sessionUser := c.Ctx.Input.Session("username")
	c.Data["sessionUser"] = sessionUser

	c.TplName = "myindex.html"
}
