package routers

import (
	"userManageSystem/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.AccountController{})
	//beego.Router("/chat", &controllers.ChatController{})
}
