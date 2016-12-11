package routers

import (
	"github.com/Baligul/humansystech/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/humansystech/view", &controllers.HumansystechController{}, "get:Form")
    beego.Router("/humansystech/view", &controllers.HumansystechController{}, "post:View")
	beego.Router("/humansystech/home", &controllers.HumansystechController{}, "*:Home")    
}
