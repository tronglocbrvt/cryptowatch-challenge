package services

import (
	"context"

	pb "github.com/cryptowatch_challenge/pb/proto"
)

func (s *service) OpenSyntheticPosition(ctx context.Context, req *pb.OpenSyntheticPositionRequest) (*pb.OpenSyntheticPositionResponse, error) {
	resp, err := s.openSyntheticPosition(ctx, req)
	return resp, err
}
