package router

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "go_dev/SecKill/SecProxy/controller"
)

func init() {
    logs.Debug("Init beego router.")
    beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
    beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}