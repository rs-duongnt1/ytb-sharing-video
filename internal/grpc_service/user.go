package grpcservice

import (
	"context"
	"fmt"
	"net/mail"
	"strings"
	"time"

	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	"github.com/youtube_sharing/youtube_sharing/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *BuildoriumGrpcService) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {

	password, err := s.Password.GeneratePassword([]byte(req.Password))

	if err != nil {
		return nil, err
	}

	_, err = mail.ParseAddress(req.Email)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(req.Password) != strings.TrimSpace(req.ConfirmPassword) {
		return nil, status.Error(codes.InvalidArgument, "Password is not match")
	}

	user := models.User{
		Email:    req.Email,
		Password: string(password),
	}

	err = s.DB.Create(&user).Error

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, _, err := s.JwtMaker.CreateToken(user.ID, user.Email, time.Hour*1)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.UserRegisterResponse{
		Id:    s.ID.Encoder(int(user.ID)),
		Email: user.Email,
		Token: token,
	}, nil
}

func (s *BuildoriumGrpcService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	var user *models.User
	err := s.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	isCompare := s.Password.ComparePassword([]byte(user.Password), []byte(req.Password))

	if !isCompare {
		return nil, fmt.Errorf("credential failed")
	}

	token, _, err := s.JwtMaker.CreateToken(user.ID, user.Email, time.Hour*1)
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginResponse{
		Token: token,
	}, nil
}

func (s *BuildoriumGrpcService) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	authHeaderSplitted := strings.Split(authHeader[0], " ")

	if len(authHeaderSplitted) != 2 {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	token := authHeaderSplitted[1]

	user, err := s.JwtMaker.VerifyToken(token)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	return &pb.UserInfoResponse{
		Id:    user.ID,
		Email: user.Email,
	}, nil
}
