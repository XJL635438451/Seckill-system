package service

import (
	"MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/common"
	"github.com/astaxie/beego/logs"
    "fmt"
    "time"
)

var (
    secKillConf *common.SecKillConf
)

func InitService(serviceConf *common.SecKillConf) {
    secKillConf = serviceConf
    logs.Debug("Init service succ, config: %v", secKillConf)
}

func SecInfoList() (data []map[string]interface{}, code int, err error) {
    secKillConf.RWSecProductLock.RLock()
    defer secKillConf.RWSecProductLock.RUnlock()

    for _, v := range secKillConf.SecProductInfoMap {
        item, _, err := SecInfoById(v.ProductId)
        if err != nil {
            logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
            continue
        }

        logs.Debug("get product[%d]， result[%v], all[%v] v[%v]", v.ProductId, item, secKillConf.SecProductInfoMap, v)
        data = append(data, item)
    }

    return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
    item, code, err := SecInfoById(productId)
    if err != nil {
        return 
    }

    data = append(data, item)
    return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {
    secKillConf.RWSecProductLock.RLock()
    defer secKillConf.RWSecProductLock.RUnlock()

    v, ok := secKillConf.SecProductInfoMap[productId]
    if !ok {
        code = common.ErrNotFoundProductId
        err = fmt.Errorf("Not found product_id:%d", productId)
        return
    }

    start := false
    end := false
    status := "success"

    now := time.Now().Unix()
    if now - v.StartTime < 0 {
        start = false
        end = false
        status = "Sec kill do not start."
        code = common.ErrActiveNotStart
    }

    if now - v.StartTime >= 0 {
        start = true
    }

    if now - v.EndTime > 0 {
        start = false
        end = true
        status = "Sec kill is already end."
        code = common.ErrActiveAlreadyEnd
    }

    if v.Status == common.ProductStatusForceSaleOut || v.Status == common.ProductStatusSaleOut {
        start = false
        end = true
        status = "Product is sale out."
        code = common.ErrActiveSaleOut
    }


    data = make(map[string]interface{})
    data["product_id"] = productId
    data["start_time"] = start
    data["end_time"] = end
    data["status"] = status

    return
}
