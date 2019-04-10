package main

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    _ "go_dev/SecKill/SecProxy/router"
)

func main() {
    err := InitSec()
    if err != nil {
        logs.Error(err)
    }
    logs.Debug("init all success")

    beego.Run()
}