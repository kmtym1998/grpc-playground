package main

import (
	"fmt"
	"grpc-playground/api/generated"
	"grpc-playground/api/rpc"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. 8080番portのLisnterを作成
	port := func() string {
		v := os.Getenv("PORT")
		if v == "" {
			return "8080"
		}
		return v
	}()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	srv := grpc.NewServer()

	generated.RegisterGreetingServiceServer(
		srv,
		rpc.NewGreetingService(),
	)

	reflection.Register(srv)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		srv.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	srv.GracefulStop()
}
