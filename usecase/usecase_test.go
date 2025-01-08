package usecase

import (
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/models"
)

var UseCaseTestConfig config.Config

func init() {
	UseCaseTestConfig = config.NewConfig("../toml/local.toml", "../env/.env.local")

	models.Init(UseCaseTestConfig)
}
