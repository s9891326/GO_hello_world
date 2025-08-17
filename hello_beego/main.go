package main

import (
	_ "hello_beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

