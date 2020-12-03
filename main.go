package main

import (
	"beegodemo01/models"
	_ "beegodemo01/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 自定义模板函数
	beego.AddFuncMap("hi", models.Hello)

	// 把时间戳转换成时间格式
	beego.AddFuncMap("unixToDate", models.UnixToDate)

	// 把时间格式化的字符串转换成时间戳
	beego.AddFuncMap("dateToUnix", models.DateToUnix)

	// 加载其他配置文件
	beego.LoadAppConfig("ini", "conf/redis.conf")

	// 启动 session
	// beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
