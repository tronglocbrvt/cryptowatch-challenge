package services

import (
	"context"
	"errors"

	"github.com/cryptowatch_challenge/internal/constants"
	"github.com/cryptowatch_challenge/internal/models"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"go.uber.org/zap"
)

func (s *service) openSyntheticPosition(ctx context.Context, req *pb.OpenSyntheticPositionRequest) (*pb.OpenSyntheticPositionResponse, error) {
	position := models.FromPositionProto(req.Position)
	if position == nil {
		s.log.Error("request is wrong format")
		return nil, errors.New("request is wrong format")
	}
	err := s.positionStore.Save(position).Error
	if err != nil {
		s.log.Error("error when opening synthetic position", zap.Error(err))
		return nil, err
	}

	return &pb.OpenSyntheticPositionResponse{
		Flag:    constants.FlagSuccess,
		Message: "Success",
	}, nil
}
