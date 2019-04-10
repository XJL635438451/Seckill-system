package common

import "sync"

//all seckill conf info
type SecKillConf struct {
    RedisConf RedisConf
    EtcdConf EtcdConf
    LogConf LogConf
    //By commodity ID to distinguish the goods
    SecProductInfoMap map[int]*SecProductInfoConf
    RWSecProductLock  sync.RWMutex
}

//log conf
type LogConf struct {
    LogPath string
    LogLevel string
}

//redis conf
type RedisConf struct {
    RedisAddr string
    RedisMaxIdle int
    RedisMaxActive int
    RedisIdleTimeout int
}

//etcd conf
type EtcdConf struct {
    EtcdAddr string
    EtcdTimeout int
    EtcdSecKeyPrefix string
    EtcdSecProductKey string
}

//sec product info
type SecProductInfoConf struct {
    ProductId int
    StartTime int
    EndTime   int
    Status    int
    Total     int
    Left      int
}
