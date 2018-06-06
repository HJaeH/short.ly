package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/short.ly/controllers"
	"github.com/short.ly/db/redis"
)

func main() {
	log := logs.NewLogger()

	log.SetLogger(logs.AdapterConsole)
	beego.SetLevel(beego.LevelInformational)
	//beego.SetLogger("file", `{"filename":"bin/shortly.log"}`)

	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/create", &controllers.MainController{}, "post:ShortURL")
	beego.Router("/?:short", &controllers.MainController{}, "get:RedirectToOriginal")

	var initialShortURL uint8 = 3
	go func() {
		redis.InitNumbers(initialShortURL)
	}()
	beego.Run()
}
