package services

import (
	"context"

	pb "github.com/cryptowatch_challenge/pb/proto"
)

func (s *service) GetLatestPrice(ctx context.Context, req *pb.GetLatestPriceRequest) (*pb.GetLatestPriceResponse, error) {
	resp, err := s.getLatestPrice(ctx, req)
	return resp, err
}

func (s *service) GetPrices(ctx context.Context, req *pb.GetPricesRequest) (*pb.GetPricesResponse, error) {
	resp, err := s.getPrices(ctx, req)
	return resp, err
}
