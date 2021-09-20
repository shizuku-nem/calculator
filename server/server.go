package main

import (
	"context"
	"fmt"
	"log"
	"main/calculator/calculatorpb"
	"net"

	"google.golang.org/grpc"
)

type server struct {}

// implement interface ServiceServer from pb.go
// request and response
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatal("err while listening %v", err)
	}

	s := grpc.NewServer() // create server without regiter

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("calculator is running")

	err = s.Serve(lis) // run server

	if err != nil {
		log.Fatal("err while serving %v", err)
	}
}