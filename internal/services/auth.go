package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cryptowatch_challenge/internal/constants"
	"github.com/cryptowatch_challenge/internal/models"
	pb "github.com/cryptowatch_challenge/pb/proto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

const expiration = 240

func (s *service) authenticationGoogle(ctx context.Context, req *pb.AuthenticationGoogleRequest) (*pb.AuthenticationGoogleResponse, error) {
	tokenInfo, err := s.getEmailIDToken(ctx, req.IdToken)
	if err != nil || tokenInfo == nil {
		s.log.Error("error when getting email from id_token", zap.Error(err))
		return nil, err
	}

	var userID uint32
	user, exist, err := s.userStore.GetByEmail(tokenInfo.Email)
	if err != nil {
		s.log.Error("error when getting user by email", zap.Error(err))
		return nil, err
	}
	if !exist {
		newUser := &models.User{
			GoogleID: tokenInfo.UserId,
			Email:    tokenInfo.Email,
		}
		err := s.userStore.Save(newUser).Error
		if err != nil {
			s.log.Error("error when creating new user", zap.Error(err))
			return nil, err
		}
		userID = newUser.UserID
	}
	userID = user.UserID

	accessToken, expireAccess, err := generateToken([]byte("SECRET_KEY_ACCESS_JWT"), fmt.Sprintf("%d", userID))
	if err != nil {
		s.log.Error("error when generating token", zap.Error(err))
		return nil, errors.New(err.Error())
	}
	refreshToken, expireRefresh, err := generateRefreshToken([]byte("SECRET_KEY_REFRESH_JWT"), fmt.Sprintf("%d", userID))
	if err != nil {
		s.log.Error("error when generating refresh token", zap.Error(err))
		return nil, errors.New(err.Error())
	}

	err = s.userStore.Update("refresh_token", refreshToken).Error
	if err != nil {
		s.log.Error("error when updating refresh token", zap.Error(err))
		return nil, errors.New(err.Error())
	}

	return &pb.AuthenticationGoogleResponse{
		Flag:        constants.FlagSuccess,
		Message:     "Success",
		AccessToken: accessToken,
		ExpAccess:   expireAccess,
		RefToken:    refreshToken,
		ExpRef:      expireRefresh,
		Email:       tokenInfo.Email,
	}, nil
}

func (s *service) regenerateAccessToken(ctx context.Context, req *pb.RegenerateAccessTokenRequest) (*pb.RegenerateAccessTokenResponse, error) {
	user, exist, err := s.userStore.GetByID(req.UserId)
	if err != nil {
		s.log.Error("error when getting user", zap.Error(err))
		return nil, err
	}
	if !exist {
		s.log.Error("user does not exist", zap.Error(err))
		return nil, err
	}

	if user.RefreshToken != req.RefreshToken {
		s.log.Error("refresh token is invalid", zap.Error(err))
		return nil, err
	}

	kf := func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY_REFRESH_JWT")), nil
	}

	claims, err := s.ExtractClaims(user.RefreshToken, kf)
	if err != nil {
		s.log.Error("error when extracting claims", zap.Error(err))
		return nil, err
	}

	accessToken, expireAccess, err := generateToken([]byte("SECRET_KEY_ACCESS_JWT"), claims["sub"].(string))
	if err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return &pb.RegenerateAccessTokenResponse{
		Flag:        constants.FlagSuccess,
		Message:     "Success",
		AccessToken: accessToken,
		ExpAccess:   expireAccess,
	}, nil
}

func (s *service) logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	kf := func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY_ACCESS_JWT"), nil
	}

	claims, err := s.ExtractClaims(req.AccessToken, kf)
	if err != nil {
		return nil, err
	}

	if fmt.Sprintf("%d", req.UserId) != claims["sub"].(string) {
		s.log.Error("access token is invalid", zap.Error(err))
		return nil, err
	}

	err = s.userStore.Update("refresh_token", "").Error
	if err != nil {
		s.log.Error("error when removing refresh_token", zap.Error(err))
		return nil, err
	}
	return &pb.LogoutResponse{
		Flag:    constants.FlagSuccess,
		Message: "Success",
	}, nil
}

func (s *service) getEmailIDToken(ctx context.Context, idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(http.DefaultClient))
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func generateToken(signingKey []byte, u string) (string, int64, error) {
	claims :=
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Subject:   u,
			Id:        uuid.New().String(),
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, claims.ExpiresAt, err
}

func generateRefreshToken(signingKey []byte, uuid string) (string, int64, error) {
	claims :=
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * expiration).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Subject:   uuid,
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, claims.ExpiresAt, err
}
