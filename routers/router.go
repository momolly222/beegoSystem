package routers

import (
	"beegodemo01/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/add", &controllers.GoodsController{}, "post:DoAdd")
	beego.Router("/goods/edit", &controllers.GoodsController{}, "put:DoEdit")
	beego.Router("/goods/delete", &controllers.GoodsController{}, "delete:DoDelete")
	beego.Router("/goods/xml", &controllers.GoodsController{}, "post:Xml")

	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle") // "get："后不能有空格
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")

	beego.Router("/user", &controllers.UserController{}) // 获取用户中心页面
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoUser")
	beego.Router("/user/getuser", &controllers.UserController{}, "get:GetUser") // 获取单个用户详细页面

	beego.Router("/api/:id", &controllers.ApiController{})
	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/loginOut", &controllers.LoginController{}, "get:DoLoginOut")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")

	beego.Router("/index", &controllers.ViewController{})
}
