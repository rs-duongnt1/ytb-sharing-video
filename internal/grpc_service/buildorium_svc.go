package grpcservice

import (
	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	"github.com/youtube_sharing/youtube_sharing/pkg/id"
	"github.com/youtube_sharing/youtube_sharing/pkg/password"
	"github.com/youtube_sharing/youtube_sharing/pkg/token"
	"gorm.io/gorm"
)

var EXECUTOR_NAMESPACE = "executor-runner"

type BuildoriumGrpcService struct {
	pb.UnimplementedBuildoriumServiceServer
	DB       *gorm.DB
	ID       *id.ID
	Password *password.Password
	JwtMaker *token.JWTMaker
}
