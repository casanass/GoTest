package main

import (
	"DEMO/models"
	"DEMO/routers"
)

func main() {
	//初始化数据库
	routers.InitRouter()
	//初始化
	models.InitDB()
}
