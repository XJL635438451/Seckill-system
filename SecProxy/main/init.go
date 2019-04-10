package main

import (
    "github.com/garyburd/redigo/redis"
    etcd_client "go.etcd.io/etcd/clientv3"
    "time"
    "fmt"
)

var (
    pool *redis.Pool
    cli *etcd_client.Client
)

func InitRedis() {
    pool = &redis.Pool {
        MaxIdle: secKillConf.redisConf.RedisMaxIdle,
        MaxActive:   secKillConf.redisConf.RedisMaxActive, 
        IdleTimeout: time.Duration(secKillConf.redisConf.RedisIdleTimeout)*time.Second,
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", secKillConf.redisConf.RedisAddr)
        },
    }
}

func InitEtcd() (err error) {
    cli, err = etcd_client.New(etcd_client.Config{
        Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
        DialTimeout: time.Duration(secKillConf.etcdConf.etcdTimeout) * time.Second,
    })
    if err != nil {
        err = fmt.Errorf("connect failed, Error: %v", err)
        return
    }
    return 
}

func InitSec() (err error) {
    InitRedis()
    err = InitEtcd()
    if err != nil {
        return 
    }
    return 
}