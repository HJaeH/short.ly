package controllers

import (
	"fmt"

	"strings"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/short.ly/models"
	"github.com/short.ly/utils/result_code"
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

	req := ShortUrlInput{}

	if err := o.ParseForm(&req); err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("Internal Error"))
		return
	}

	if req.URL == "" {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte(resultcode.ResultCodeMap[resultcode.ErrorURLIsRequired]))
		return
	}

	shortURL, err := models.AddURL(req.URL)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte(" "))
		return
	}
	beego.Info("original URL [", req.URL, "], => short URL [", shortURL, "]")
	result := models.URL{
		OriginalURL: req.URL,
		ShortURL:    shortURL,
	}
	byteResult, err := json.Marshal(result)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		o.Ctx.Output.Body([]byte("Internal Error"))
		return
	}
	o.Ctx.Output.Body([]byte(byteResult))
}

func (o *MainController) RedirectToOriginal() {

	fmt.Println("RedirectToOriginal ")
	shortUrl := strings.TrimPrefix(o.Ctx.Request.RequestURI, "/")
	originalUrl, err := models.GetOriginalUrl(shortUrl)
	if err != nil {
		o.Ctx.Output.SetStatus(404)
	}

	fmt.Println("short url : ", shortUrl, ", original url :", originalUrl)
	o.Ctx.Output.Body([]byte(originalUrl))
	//models.GetUrl()

	//o.Redirect()
}
