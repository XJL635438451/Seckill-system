package main

import (
    "time"
    "fmt"
    "strings"
    "encoding/json"
    "context"
    "go.etcd.io/etcd/mvcc/mvccpb"
    "github.com/astaxie/beego/logs"
    "github.com/garyburd/redigo/redis"
    etcd_client "go.etcd.io/etcd/clientv3"
    "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/common"
    "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/service"
)

var (
    redisPool *redis.Pool
    etcdClient *etcd_client.Client
)

// init logger
func convertLogLevel(level string) int {
    level = strings.ToLower(level)
    switch (level) {
    case "debug":
        return logs.LevelDebug
    case "warn":
        return logs.LevelWarn
    case "info":
        return logs.LevelInfo
    case "trace":
        return logs.LevelTrace
    }

    return  logs.LevelDebug
}

func initLogger() (err error) {
    logs.Debug("Starting to init logger.")
    config := make(map[string]interface{})
    config["filename"] = secKillConf.LogConf.LogPath
    config["level"] = convertLogLevel(secKillConf.LogConf.LogLevel)

    configStr, err := json.Marshal(config)
    if err != nil {
        err = fmt.Errorf("Init logs failed, Error: %v", err)
        return
    }
    logs.Debug("Successfully initialized logger.")
    logs.SetLogger(logs.AdapterFile, string(configStr))
    return
}

//init redis
func initRedis() (err error) {
    logs.Debug("Starting to init redis.")
    redisPool = &redis.Pool {
        MaxIdle: secKillConf.RedisConf.RedisMaxIdle,
        MaxActive:   secKillConf.RedisConf.RedisMaxActive, 
        IdleTimeout: time.Duration(secKillConf.RedisConf.RedisIdleTimeout)*time.Second,
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", secKillConf.RedisConf.RedisAddr)
        },
    }

    conn := redisPool.Get()
    defer conn.Close()

    _, err = conn.Do("ping")
    if err != nil {
        logs.Error("Ping redis failed, Error: %v", err)
        return 
    }
    logs.Debug("Successfully initialized redis.")
    return
}

//init etcd
func initEtcd() (err error) {
    logs.Debug("Starting to init etcd.")
    client, err := etcd_client.New(etcd_client.Config{
        Endpoints:   []string{secKillConf.EtcdConf.EtcdAddr},
        DialTimeout: time.Duration(secKillConf.EtcdConf.EtcdTimeout) * time.Second,
    })
    if err != nil {
        err = fmt.Errorf("Connect etcd server failed, Error: %v", err)
        return
    }
    etcdClient = client
    logs.Debug("Successfully initialized etcd.")
    return 
}

func loadSecConf() (err error) {
    logs.Debug("Starting to load sec config.")
    resp, err := etcdClient.Get(context.Background(), secKillConf.EtcdConf.EtcdSecProductKey)
    if err != nil {
        logs.Error("Get [%s] from etcd failed, Error: %v", secKillConf.EtcdConf.EtcdSecProductKey, err)
        return
    }

    var secProductInfo []common.SecProductInfoConf

    for k, v := range resp.Kvs {
        logs.Debug("key[%v] valud[%v]", k, v)
        err = json.Unmarshal(v.Value, &secProductInfo)
        if err != nil {
            logs.Error("Unmarshal sec product info failed, Error: %v", err)
            return
        }

        logs.Debug("Get sec conf info success. it is [%v]", secProductInfo)
    }

    updateSecProductInfo(secProductInfo)
    logs.Debug("Successfully load sec config.")
    return
}

func updateSecProductInfo(secProductInfo []common.SecProductInfoConf) {
    var tmpProductInfoMap map[int]*common.SecProductInfoConf = make(map[int]*common.SecProductInfoConf, 1024)
    for _, v := range secProductInfo {
        product := v
        tmpProductInfoMap[v.ProductId] = &product
        //tmpProductInfoMap[v.ProductId] = &product //this is a bug
        //example: secProductInfo -> [{1029 1505008800 1505012400 0 1000 1000} {1027 1505008800 1505012400 0 2000 1000}]
        //wait loop end : tmpProductInfoMap -> map[1027:0xc000180840 1029:0xc000180840], 1027 addr is the same as 1029
        //because v is variable and its add is invariable.
        logs.Debug("Temp:%v",tmpProductInfoMap)
    }

    secKillConf.RWSecProductLock.Lock()
    secKillConf.SecProductInfoMap = tmpProductInfoMap
    secKillConf.RWSecProductLock.Unlock()
}

func initSecProductWatcher() {
    go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
}

func watchSecProductKey(key string) {
    cli, err := etcd_client.New(etcd_client.Config{
        Endpoints:   []string{secKillConf.EtcdConf.EtcdAddr},
        DialTimeout: time.Duration(secKillConf.EtcdConf.EtcdTimeout) * time.Second,
    })
    if err != nil {
        logs.Error("Connect etcd failed, err:", err)
        return
    }

    logs.Debug("Begin watch key: %s", key)
    for {
        rch := cli.Watch(context.Background(), key)
        var secProductInfo []common.SecProductInfoConf
        var getConfSucc = true

        for wresp := range rch {
            for _, ev := range wresp.Events {
                if ev.Type == mvccpb.DELETE {
                    logs.Warn("key[%s] 's config deleted", key)
                    continue
                }

                if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
                    err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
                    if err != nil {
                        logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
                        getConfSucc = false
                        continue
                    }
                }
                logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
            }

            if getConfSucc {
                logs.Debug("get config from etcd succ, %v", secProductInfo)
                updateSecProductInfo(secProductInfo)
            }
        }

    }
}

func InitSec() (err error) {
    err = initLogger()
    if err != nil {
        return 
    }

    err = initRedis()
    if err != nil {
        return
    }

    err = initEtcd()
    if err != nil {
        return 
    }

    err = loadSecConf()
    if err != nil {
        return 
    }

    service.InitService(secKillConf)
    initSecProductWatcher()

    logs.Info("init sec succ")
    return 
}