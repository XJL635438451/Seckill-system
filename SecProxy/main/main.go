package main

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    _ "MyGitHubProject/SrcKillProject/Seckill-system/SecProxy/router"
)

func main() {
    err := initConfig()
    if err != nil {
        panic(err)
    }

    err = InitSec()
    if err != nil {
        logs.Error(err)
        return 
    }


    beego.Run()
}