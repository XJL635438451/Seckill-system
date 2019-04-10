package controller

import (
    "github.com/astaxie/beego"
)

type SkillController struct {
    beego.Controller
}

func (p *SkillController) SecKill() {
    p.Data["json"] = "hello seckill"
    p.ServeJSON()
}

func (p *SkillController) SecInfo() {
    p.Data["json"] = "hello secinfo"
    p.ServeJSON()
}