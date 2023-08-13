package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SocketServer struct {
	conns  map[*websocket.Conn]bool
	client pb.BuildoriumServiceClient
}

func NewSocketServer() *SocketServer {
	conn, _ := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := pb.NewBuildoriumServiceClient(conn)
	return &SocketServer{
		conns:  make(map[*websocket.Conn]bool),
		client: client,
	}
}

func (s *SocketServer) handleWs(ws *websocket.Conn) {
	defer ws.Close()
	s.conns[ws] = true

	// remove conn on client disconnected
	defer func() {
		delete(s.conns, ws)
		ws.Close()
		fmt.Println("Client disconnected")
	}()
	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}
		var data struct {
			Event   string
			Message string
			Token   string
		}
		err = json.Unmarshal([]byte(message), &data)
		if err != nil {
			fmt.Println(err)
			break
		}
		video, err := s.client.VideoShare(context.Background(), &pb.VideoShareRequest{
			Url:   data.Message,
			Token: data.Token,
		})

		if err != nil {
			fmt.Println(err)
			websocket.Message.Send(ws, "video not found")
			continue
		}

		fmt.Println("Received:", message)
		fmt.Println(video)

		b, err := json.Marshal(video)
		s.broadcast(b)
	}

}

func (s *SocketServer) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		fmt.Println(msg)
	}
}

func (s *SocketServer) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func (s SocketServer) start() error {
	http.Handle("/ws", websocket.Handler(s.handleWs))
	log.Println("starting websocket server addr :8000")
	return http.ListenAndServe(":8000", nil)
}
