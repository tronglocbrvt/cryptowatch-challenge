package models

import (
	pb "github.com/cryptowatch_challenge/pb/proto"
)

func ToPriceProto(model *Price) *pb.Price {
	if model == nil {
		return nil
	}

	price := &pb.Price{
		PriceId:   model.PriceID,
		MarketId:  model.MarketID,
		Price:     model.Price,
		CreatedAt: model.CreatedAt.String(),
	}

	return price
}

func FromPositionProto(proto *pb.Position) *Position {
	if proto == nil {
		return nil
	}

	position := &Position{
		UserID:     proto.UserId,
		Asset:      proto.Asset,
		Side:       proto.Side,
		Size:       proto.Size,
		EntryPrice: proto.EntryPrice,
		Leverage:   proto.Leverage,
		Status:     proto.Status,
		MarketID:   proto.MarketId,
	}

	return position
}
