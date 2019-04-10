package controller

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/common"
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
    result := make(map[string]interface{})

    result["code"] = 0
    result["message"] = "success"

    //return client format is json
    defer func() {
        p.Data["json"] = result
        p.ServeJSON()
    }()

    productId, err := p.GetInt("product_id")
    if err != nil {
        result["code"] = common.ErrInvalidRequest
        result["message"] = err.Error()
        logs.Error("Invalid request, get product_id failed, Error: %v", err)
        return 
    }
    logs.Debug("Get productId success. productId: %v", productId)
   
    data, code, err := service.SecInfo(productId)
    if err != nil {
        result["code"] = code
        result["message"] = err.Error()
        logs.Error(err)
        return 
    }
    
    result["data"] = data
}