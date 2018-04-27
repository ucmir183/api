package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

const (
	MSG_OK=0
	MSG_ERR = -1
)


func ( this *BaseController ) ToData(msg interface{},code int,data interface{}) {
	t := make(map[string]interface{})
	t["msg"] = msg
	t["status"] = code
	t["data"] = data

	this.Data["json"] = t
	this.ServeJSON()
	this.StopRun()
}

func ( this *BaseController ) ToMsg(msg interface{},code int) {
	t := make(map[string]interface{})
	t["msg"] = msg
	t["status"] = code

	this.Data["json"] = t
	this.ServeJSON()
	this.StopRun()
}