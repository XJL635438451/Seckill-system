package controller

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/service"
)

type SkillController struct {
    beego.Controller
}

func (p *SkillController) SecKill() {
    p.Data["json"] = "hello seckill"
    p.ServeJSON()
}

func (p *SkillController) SecInfo() {
    productId, err := p.GetInt("product_id")
    result := make(map[string]interface{})

    result["code"] = 0
    result["message"] = "success"

    defer func() {
        p.Data["json"] = result
        p.ServeJSON()
    }()

    if err != nil {
        //Not get product_id, list all sec product message
        data, code, err := service.SecInfoList()
        if err != nil {
            result["code"] = code
            result["message"] = err.Error()

            logs.Error("Invalid request, get product_id failed, err:%v", err)
            return
        }
        result["code"] = code
        result["data"] = data
    } else {
        //product_id exist, list the sec product message
        data, code, err := service.SecInfo(productId)
        if err != nil {
            result["code"] = code
            result["message"] = err.Error()

            logs.Error("Get product_id success, but get secinfo failed. err:%v", err)
            return
        }
        result["code"] = code
        result["data"] = data
    }
}
