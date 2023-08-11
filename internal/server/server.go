package server

import (
	"flag"
	"log"

	"github.com/youtube_sharing/youtube_sharing/internal/models"
	"github.com/youtube_sharing/youtube_sharing/pkg/config"
	"github.com/youtube_sharing/youtube_sharing/pkg/db"
	"github.com/youtube_sharing/youtube_sharing/pkg/id"
	"github.com/youtube_sharing/youtube_sharing/pkg/password"
	"github.com/youtube_sharing/youtube_sharing/pkg/token"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func Run() {
	flag.Parse()
	config.InitConfig()

	fx.New(
		fx.Options(
			fx.Provide(newGrpcServer),
			fx.Provide(NewSocketServer),
			fx.Provide(newHttpServer),
			fx.Provide(db.NewPgSQL),
			fx.Provide(password.NewPassword),
			fx.Provide(token.NewJWTMaker),
			fx.Provide(id.NewID),

			fx.Invoke(start),
		),
	).Run()

}

func start(
	httpServer *httpServer,
	grpcServer *grpcServer,
	socketServer *SocketServer,
	db *gorm.DB,
) {
	autoMigrate(db)
	go func() {
		err := grpcServer.start()
		if err != nil {
			log.Fatalln("failed running the gRPC server", "err", err)
		}
	}()

	go func() {
		err := socketServer.start()
		if err != nil {
			log.Fatalln("failed running the Socket server", "err", err)
		}
	}()

	err := httpServer.start()
	if err != nil {
		log.Fatalln("failed running the HTTP server", "err", err)
	}

}

func autoMigrate(
	db *gorm.DB,
) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Video{})
}
