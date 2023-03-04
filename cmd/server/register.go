package main

import (
	"github.com/cryptowatch_challenge/config"
	"github.com/cryptowatch_challenge/external"
	"github.com/cryptowatch_challenge/internal/services"
	"github.com/cryptowatch_challenge/internal/stores"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

func registerService(cfg *config.Config, log *zap.Logger) pb.CryptoWatchServer {
	db := mustConnectPostgreSQL(cfg)

	cryptoWatchClient := external.NewCryptoWatchClient(cfg)

	priceStore := stores.NewPriceStore(db)
	userStore := stores.NewUserStore(db)
	positionStore := stores.NewPositionStore(db)

	return services.New(cfg, log, cryptoWatchClient, priceStore, userStore, positionStore)
}
