package services

import (
	"context"

	pb "github.com/cryptowatch_challenge/pb/proto"
)

func (s *service) echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{
		Value: req.Value,
	}, nil
}
