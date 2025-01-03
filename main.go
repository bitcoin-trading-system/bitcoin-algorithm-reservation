package main

import (
	"flag"
	"net/http"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/db"
	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
)

func main() {
	tomlFilePath := flag.String("conf", "toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	db.ConnectDB(cfg)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	if err := router.Run(":8003"); err != nil {
		panic(err)
	}
}
