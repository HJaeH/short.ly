package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/my-beego-todo/controllers"
	"./utils"
)

func main() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("this is a debug message")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:id:int", &controllers.MainController{}, "get:Redirect")
	beego.Router("/shorten", &controllers.MainController{}, "post:ShortURL")
	beego.Run()

	utils.InitCounter()

}
