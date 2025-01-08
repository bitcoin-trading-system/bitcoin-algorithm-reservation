package main

import (
	"flag"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/models"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/router"
)

func main() {
	tomlFilePath := flag.String("conf", "toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	if err := models.Init(cfg); err != nil {
		panic(err)
	}

	router := router.NewRouter(cfg)

	if err := router.Run(":8003"); err != nil {
		panic(err)
	}
}
