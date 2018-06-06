package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/short.ly/models"
)

type ShortUrlInput struct {
	URL string `form:"url"`
}

type MainController struct {
	beego.Controller
}

func (o *MainController) Get() {
	o.TplName = "index.html"
	o.Render()
}

func (o *MainController) ShortURL() {
	beego.Debug("ShortURL")

	req := ShortUrlInput{}
	fmt.Println("ShortURL , url :", o.Input().Get("url"), o.Ctx.Input.Param(":url"))

	if err := o.ParseForm(&req); err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("Internal Error"))
		return
	}

	if req.URL == "" {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("URL Empty "))
		return
	}

	models.AddURL(req.URL)
	o.Ctx.Output.Body([]byte("ok"))
}

func (o *MainController) RedirectToOriginal() {

	fmt.Println("RedirectToOriginal")
	req := ShortUrlInput{}

	if err := o.ParseForm(&req); err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("Internal Error"))
		return
	}
	//models.GetUrl()

	//o.Redirect()
}
