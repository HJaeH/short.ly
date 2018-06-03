package controllers

import (
	"encoding/json"

	"fmt"

	"github.com/astaxie/beego"
	"../models"
)

type MainController struct {
	beego.Controller
}

func (o *MainController) Get() {
	o.TplName = "index.html"
	o.Render()
}

func (o *MainController) ShortURL() {
	req := struct{ RawURL string }{}
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &req); err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("empty title"))
		return
	}

	t, err := models.NewURL(req.RawURL)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	fmt.Println("--", req.RawURL)
	models.DefaultURLMap.AddURL(t)
}
