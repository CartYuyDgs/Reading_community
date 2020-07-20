package main

import (
	_ "Reading_community/mbook/routers"
	_ "Reading_community/mbook/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
