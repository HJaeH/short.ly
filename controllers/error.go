package controllers

//
//import (
//	"github.com/astaxie/beego"
//)
//
//type ErrorController struct {
//	beego.Controller
//}
//
//func (c *ErrorController) Error404() {
//	c.Data["content"] = "page not found"
//	c.TplName = "404.tpl"
//}
//
//func (c *ErrorController) Error500() {
//	c.Data["content"] = "internal server error"
//	c.TplName = "500.tpl"
//}
//
//func (c *ErrorController) ErrorDb() {
//	c.Data["content"] = "database is now down"
//	c.TplName = "dberror.tpl"
//}