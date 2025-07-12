package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"grpc-multiplex/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedGreetServiceServer
}

func (s *server) Greet(req *proto.GreetRequest, stream proto.GreetService_GreetServer) error {
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("Hola %s! Respuesta %d", req.Name, i)
		if err := stream.Send(&proto.GreetResponse{Message: msg}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &server{})

	// 👇 habilitar reflection para grpcurl
	reflection.Register(grpcServer)

	log.Println("Servidor gRPC escuchando en el puerto 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
