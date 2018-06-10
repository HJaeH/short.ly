package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/HJaeH/short.ly/controllers"
	"github.com/HJaeH/short.ly/db/redis"
)

func main() {
	log := logs.NewLogger()

	log.SetLogger(logs.AdapterConsole)
	//beego.SetLogger("file", `{"filename":"bin/shortly.log"}`)
	beego.Router("/create", &controllers.MainController{}, "post:ShortURL")
	beego.Router("/list", &controllers.MainController{}, "get:GetUrlList")
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/:short", &controllers.MainController{}, "get:RedirectToOriginal")

	go func() {
		redis.ExpendCount()
	}()
	beego.Run()
}
