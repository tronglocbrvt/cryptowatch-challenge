package main

import (
	"github.com/cryptowatch_challenge/config"
	"github.com/cryptowatch_challenge/internal/services"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

func registerService(cfg *config.Config, log *zap.Logger) pb.CryptoWatchServer {

	return services.New(cfg, log)
}
