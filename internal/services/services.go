package services

import (
	"github.com/cryptowatch_challenge/config"
	"github.com/cryptowatch_challenge/external"
	"github.com/cryptowatch_challenge/internal/stores"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

const Version = "1.0.0"

type service struct {
	cfg               *config.Config
	log               *zap.Logger
	cryptoWatchClient *external.CryptoWatchClient
	priceStore        *stores.PriceStore
	userStore         *stores.UserStore
	positionStore     *stores.PositionStore

	pb.UnimplementedCryptoWatchServer
}

func New(config *config.Config,
	log *zap.Logger,
	cryptoWatchClient *external.CryptoWatchClient,
	priceStore *stores.PriceStore,
	userStore *stores.UserStore,
	positionStore *stores.PositionStore) pb.CryptoWatchServer {
	s := &service{
		cfg:               config,
		log:               log,
		cryptoWatchClient: cryptoWatchClient,
		priceStore:        priceStore,
		userStore:         userStore,
		positionStore:     positionStore,
	}
	return s
}
