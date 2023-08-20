package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/ldmtam/calculator_service/proto"

	"google.golang.org/grpc"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func newCalculatorServer() *calculatorServer {
	return &calculatorServer{}
}

func (s *calculatorServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Printf("Received Add request. A: %v, B: %v\n", in.A, in.B)
	return &pb.AddResponse{Result: in.A + in.B}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	fmt.Printf("Received Subtract request. A: %v, B: %v\n", in.A, in.B)
	return &pb.SubtractResponse{Result: in.A - in.B}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	fmt.Printf("Received Multiply request. A: %v, B: %v\n", in.A, in.B)
	return &pb.MultiplyResponse{Result: in.A * in.B}, nil
}

func (s *calculatorServer) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideResponse, error) {
	fmt.Printf("Received Divide request. A: %v, B: %v\n", in.A, in.B)
	if in.B == 0 {
		return nil, errors.New("can not divide by 0")
	}

	return &pb.DivideResponse{Result: in.A / in.B}, nil
}

func (s *calculatorServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	fmt.Println("Received Ping request")
	return &pb.PingResponse{Message: "pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, newCalculatorServer())
	grpcServer.Serve(lis)
}
