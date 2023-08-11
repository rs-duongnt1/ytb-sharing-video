package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type httpServer struct {
}

func newHttpServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) start() error {

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial gRPC server:", err)
	}
	gwmux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		token, err := r.Cookie("token")
		if err != nil {
			fmt.Println(err)
			return metadata.Pairs()
		}
		return metadata.Pairs("token", token.Value)
	}))
	gwErr := pb.RegisterBuildoriumServiceHandler(context.Background(), gwmux, conn)
	if gwErr != nil {
		log.Fatalln("Failed to register handler:", err)
	}
	// gwmux.HandlePath()
	gwServer := &http.Server{
		Addr:    ":" + "8090",
		Handler: cors(gwmux),
	}
	log.Println("starting HTTP server", "addr", ":8090")
	return gwServer.ListenAndServe()
}
