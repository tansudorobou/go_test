package main

import (
	"app/routers"
	"app/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	err := utils.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Ginルーターの初期化
	router := gin.Default()

	// 静的ファイルのルーティング
	// router.Static("/", ".../client/dist")

	// ルーティングの設定
	routers.SetupRouter(router)

	router.Run(":5000")
}
