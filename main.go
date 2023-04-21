package main

import (
	_ "DEMO/models"
	_ "DEMO/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("main init")
	fmt.Println(beego.VERSION)
}
