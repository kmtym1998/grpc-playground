package rpc

import (
	"context"
	"fmt"
	"grpc-playground/api/generated"
)

type GreetingService struct {
	generated.UnimplementedGreetingServiceServer
}

func NewGreetingService() GreetingService {
	return GreetingService{}
}

func (s GreetingService) Hello(ctx context.Context, req *generated.HelloRequest) (*generated.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &generated.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
