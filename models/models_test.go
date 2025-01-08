package models

import "github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"

func init() {
	cfg := config.NewConfig("../toml/local.toml", "../env/.env.local")
	if err := Init(cfg); err != nil {
		panic(err)
	}
}
