package services

import (
	"context"

	pb "github.com/cryptowatch_challenge/pb/proto"
)

func (s *service) AuthenticationGoogle(ctx context.Context, req *pb.AuthenticationGoogleRequest) (*pb.AuthenticationGoogleResponse, error) {
	resp, err := s.authenticationGoogle(ctx, req)
	return resp, err
}

func (s *service) RegenerateAccessToken(ctx context.Context, req *pb.RegenerateAccessTokenRequest) (*pb.RegenerateAccessTokenResponse, error) {
	resp, err := s.regenerateAccessToken(ctx, req)
	return resp, err
}

func (s *service) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	resp, err := s.logout(ctx, req)
	return resp, err
}
