package services

import (
	"github.com/cryptowatch_challenge/config"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

const Version = "1.0.0"

type service struct {
	cfg *config.Config
	log *zap.Logger
	pb.UnimplementedCryptoWatchServer
}

func New(config *config.Config, log *zap.Logger) pb.CryptoWatchServer {
	s := &service{
		cfg: config,
		log: log,
	}
	return s
}
