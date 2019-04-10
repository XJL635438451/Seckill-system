package service

import (
	"MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/common"
	"github.com/astaxie/beego/logs"
    "fmt"
)

var (
    secKillConf *common.SecKillConf
)

func InitService(serviceConf *common.SecKillConf) {
    secKillConf = serviceConf
    logs.Debug("Init service succ, config: %v", secKillConf)
}


func SecInfo(productId int) (data map[string]interface{}, code int, err error) {
    secKillConf.RWSecProductLock.RLock()
    defer secKillConf.RWSecProductLock.RUnlock()

    v, ok := secKillConf.SecProductInfoMap[productId]
    if !ok {
        code = common.ErrNotFoundProductId
        err = fmt.Errorf("Not found product_id:%d", productId)
        return
    }

    data = make(map[string]interface{})
    data["product_id"] = productId
    data["start_time"] = v.StartTime
    data["end_time"] = v.EndTime
    data["status"] = v.Status

    return
}
