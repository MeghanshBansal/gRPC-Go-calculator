package main

import (
	pb "MeghanshBansal/calculator/proto"
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Server struct{
	pb.UnsafeCalculatorServer
}

// Add implements proto.CalculatorServer.
func (s *Server) Add(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{Res: in.Num1 + in.Num2}, nil
}

// Divide implements proto.CalculatorServer.
func (s *Server) Divide(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	if in.Num2 == 0 {
		return nil, status.Errorf(status.Code(errors.New("failed to perform division operation")), "num2 can not be 0")
	}
	return &pb.CalculatorResponse{Res: in.Num1 / in.Num2}, nil
}

// Multiply implements proto.CalculatorServer.
func (s *Server) Multiply(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{Res: in.Num1 * in.Num2}, nil
}

// Subtract implements proto.CalculatorServer.
func (s *Server) Subtract(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{Res: in.Num1 - in.Num2}, nil
}

func main() {
	lst, err := net.Listen("tcp", "localhost:50012")
	if err != nil {
		log.Fatalf("failed to start the connection")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lst); err != nil {
		log.Fatalf("failed to start the server")
	}
	log.Println("server started at: ", "localhost:50012")

}
