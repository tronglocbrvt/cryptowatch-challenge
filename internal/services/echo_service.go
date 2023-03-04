package services

import (
	"context"

	pb "github.com/cryptowatch_challenge/pb/proto"
)

func (s *service) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	resp, err := s.echo(ctx, req)
	return resp, err
}
