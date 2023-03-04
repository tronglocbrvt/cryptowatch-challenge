package services

import (
	"context"

	"github.com/cryptowatch_challenge/internal/constants"
	"github.com/cryptowatch_challenge/internal/models"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

func (s *service) getLatestPrice(ctx context.Context, req *pb.GetLatestPriceRequest) (*pb.GetLatestPriceResponse, error) {
	latestPrice, exist, err := s.priceStore.GetLastestPrice(req.MarketId)
	if err != nil {
		s.log.Error("error when getting latest price", zap.Error(err))
		return nil, err
	}
	if !exist {
		s.log.Error("latest price does not exist", zap.Error(err))
		return nil, err
	}

	return &pb.GetLatestPriceResponse{
		Flag:    constants.FlagSuccess,
		Message: "Success",
		Price:   models.ToPriceProto(latestPrice),
	}, nil
}

func (s *service) getPrices(ctx context.Context, req *pb.GetPricesRequest) (*pb.GetPricesResponse, error) {
	if req.Limit <= 0 {
		req.Limit = 100
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit

	prices, err := s.priceStore.GetPrices(req.MarketId, req.Limit, offset)
	if err != nil {
		s.log.Error("error when getting latest price", zap.Error(err))
		return nil, err
	}

	pbPrices := make([]*pb.Price, 0)

	for i := 0; i < len(prices); i++ {
		pbPrices = append(pbPrices, models.ToPriceProto(prices[i]))
	}

	return &pb.GetPricesResponse{
		Flag:    constants.FlagSuccess,
		Message: "Success",
		Prices:  pbPrices,
	}, nil
}
