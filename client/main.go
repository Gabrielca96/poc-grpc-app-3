package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-multiplex/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	req := &pb.GreetRequest{Name: "Gabriel"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.Greet(ctx, req)
	if err != nil {
		log.Fatalf("Error al llamar Greet: %v", err)
	}

	log.Println("Esperando respuestas del servidor...")

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("No hay más mensajes.")
			break
		}
		if err != nil {
			log.Fatalf("Error recibiendo mensaje: %v", err)
		}
		log.Printf("Mensaje del servidor: %s", res.GetMessage())
	}
}
