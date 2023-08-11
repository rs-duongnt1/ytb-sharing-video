package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	grpcservice "github.com/youtube_sharing/youtube_sharing/internal/grpc_service"
	"github.com/youtube_sharing/youtube_sharing/pkg/id"
	"github.com/youtube_sharing/youtube_sharing/pkg/password"
	"github.com/youtube_sharing/youtube_sharing/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type grpcServer struct {
	server   *grpc.Server
	db       *gorm.DB
	id       *id.ID
	password *password.Password
	jwtMaker *token.JWTMaker
}

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)
	if info.FullMethod != "/ytb_sharing.BuildoriumService/VideoList" && info.FullMethod != "/ytb_sharing.BuildoriumService/UserLogin" && info.FullMethod != "/ytb_sharing.BuildoriumService/UserRegister" && info.FullMethod != "/ytb_sharing.BuildoriumService/VideoShare" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	if token == "" {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	return nil
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func newGrpcServer(
	db *gorm.DB,
	id *id.ID,
	password *password.Password,
	jwtMaker *token.JWTMaker,
) *grpcServer {
	s := grpc.NewServer(
		withServerUnaryInterceptor(),
	)
	return &grpcServer{
		server:   s,
		db:       db,
		id:       id,
		password: password,
		jwtMaker: jwtMaker,
	}
}

func (s *grpcServer) start() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	pb.RegisterBuildoriumServiceServer(s.server, &grpcservice.BuildoriumGrpcService{
		DB:       s.db,
		ID:       s.id,
		Password: s.password,
		JwtMaker: s.jwtMaker,
	})
	log.Println("starting gRPC server", "addr", ":8080")
	return s.server.Serve(lis)
}
