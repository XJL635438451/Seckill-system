package main

import (
    "github.com/astaxie/beego"
    "fmt"
)

var (
    secKillConf = &SecKillConf{}
)

type SecKillConf struct {
    redisConf RedisConf
    etcdConf EtcdConf
    logConf LogConf
}

type RedisConf struct {
    RedisAddr string
    RedisMaxIdle int
    RedisMaxActive int
    RedisIdleTimeout int
}

type EtcdConf struct {
    etcdAddr string
    etcdTimeout int
}

type LogConf struct {
    logPath string
    logLevel string
}

func initConfig() (err error) {
    secKillConf.redisConf.RedisIdleTimeout, err = beego.AppConfig.Int("redisIdleTimeout")
    if err != nil {
        err = fmt.Errorf("Failed to init redis. Error: %v", err)
        return 
    }
    secKillConf.redisConf.RedisAddr = beego.AppConfig.String("redisAddr")
    secKillConf.redisConf.RedisMaxIdle, err  = beego.AppConfig.Int("redisMaxIdle")
    if err != nil {
        err = fmt.Errorf("Failed to init redis. Error: %v", err)
        return 
    }
    secKillConf.redisConf.RedisMaxActive, err = beego.AppConfig.Int("redisMaxActive")
    if err != nil {
        err = fmt.Errorf("Failed to init redis. Error: %v", err)
        return 
    }

    secKillConf.logConf.logLevel = beego.AppConfig.String("logLevel")
    secKillConf.logConf.logPath = beego.AppConfig.String("logPath")

    secKillConf.etcdConf.etcdAddr = beego.AppConfig.String("etcdAddr")
    secKillConf.etcdConf.etcdTimeout, err = beego.AppConfig.Int("etcdTimeout")
    if err != nil {
        err = fmt.Errorf("Failed to init etcd. Error: %v", err)
        return 
    }
    return 
}