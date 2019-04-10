package main

import (
    "fmt"
    "strings"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/common"
)

var (
    secKillConf = &common.SecKillConf{
        SecProductInfoMap: make(map[int]*common.SecProductInfoConf, 1024),
    }
)

func configLogs() (err error) {
    logs.Debug("Starting to init logs config.")
    secKillConf.LogConf.LogLevel = beego.AppConfig.String("logLevel")
    secKillConf.LogConf.LogPath = beego.AppConfig.String("logPath")
    if len(secKillConf.LogConf.LogLevel) == 0 || len(secKillConf.LogConf.LogPath) == 0 {
        err = fmt.Errorf("Init logs failed, the length of logLevel or logPath is empty.")
        return 
    }
    logs.Debug("Successfully initialized logs config. LogConf: %v", secKillConf.LogConf)
    return
}

func configRedis() (err error) {
    logs.Debug("Starting to init redis config.")
    secKillConf.RedisConf.RedisIdleTimeout, err = beego.AppConfig.Int("redisIdleTimeout")
    if err != nil {
        err = fmt.Errorf("Init redis failed, read redisIdleTimeout failed. Error: %v", err)
        return 
    }

    secKillConf.RedisConf.RedisAddr = beego.AppConfig.String("redisAddr")
    if len(secKillConf.RedisConf.RedisAddr) == 0 {
        err = fmt.Errorf("Init redis failed, read redisAddr failed. it is empty")
        return 
    }

    secKillConf.RedisConf.RedisMaxIdle, err  = beego.AppConfig.Int("redisMaxIdle")
    if err != nil {
        err = fmt.Errorf("Init redis failed, read redisMaxIdle failed. Error: %v", err)
        return 
    }

    secKillConf.RedisConf.RedisMaxActive, err = beego.AppConfig.Int("redisMaxActive")
    if err != nil {
        err = fmt.Errorf("Init redis failed, read redisMaxActive failed. Error: %v", err)
        return 
    }
    logs.Debug("Successfully initialized redis config. RedisConf: %v.", secKillConf.RedisConf)
    return 
}

func configEtcd()  (err error) {
    logs.Debug("Starting to init etcd config.")
    secKillConf.EtcdConf.EtcdAddr = beego.AppConfig.String("etcdAddr")
    if len(secKillConf.EtcdConf.EtcdAddr) == 0 {
        err = fmt.Errorf("Init etcd failed, read etcdAddr failed. it is empty.")
        return 
    }

    secKillConf.EtcdConf.EtcdTimeout, err = beego.AppConfig.Int("etcdTimeout")
    if err != nil {
        err = fmt.Errorf("Init etcd failed, read etcdTimeout failed. Error: %v", err)
        return 
    }

    secKillConf.EtcdConf.EtcdSecKeyPrefix = beego.AppConfig.String("etcdSecKeyPrefix")
    if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0 {
        err = fmt.Errorf("Init etcd failed, read etcdSecKeyPrefix failed. it is empty.")
        return 
    }
    if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix, "/") == false {
        secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix + "/"
    }

    productKet := beego.AppConfig.String("etcdProductKey")
    if len(productKet) == 0 {
        err = fmt.Errorf("Init config failed, read etcdProductKey failed. error:%v", err)
        return
    }
    secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s", secKillConf.EtcdConf.EtcdSecKeyPrefix ,productKet)
    if len(secKillConf.EtcdConf.EtcdSecProductKey) == 0 {
        err = fmt.Errorf("Init etcd failed, read etcdSecProductKey failed. it is empty.")
        return 
    }

    logs.Debug("Successfully initialized etcd config. EtcdConf: %v.", secKillConf.EtcdConf)
    return
}

func initConfig() (err error) {
    err = configLogs()
    if err != nil {
        return 
    }

    err = configRedis()
    if err != nil {
        return 
    }
    
    err = configEtcd()
    if err != nil {
        return 
    }

    logs.Debug("Init all config success. secKillConf: %v", secKillConf)
    return 
}