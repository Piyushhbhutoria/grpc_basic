package main

import (
	"context"
	"grpc_tut/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":4040"
)

type server struct{}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request)(*proto.Response,error ){
	a,b :=request.GetA(),request.GetB()
	result :=a+b
	return &proto.Response{Result:result},nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request)(*proto.Response,error ){
	a,b :=request.GetA(),request.GetB()
	result :=a*b
	return &proto.Response{Result:result},nil
}
