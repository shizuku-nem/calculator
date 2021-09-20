package main

import (
	"context"
	"log"
	"main/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure()) // note: WithInsecure is no secure
	// often: grpc suggest xu dung ssl secure

	if err != nil {
		log.Fatal("err while dial %v", err)
	}

	// close dial connection
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	log.Printf("service client %f", client)

	// client.<name method> to call method
	callSum(client)
}

// input: clientService from pb
func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Call sum api grpc...")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 6,
	})

	if err != nil {
		log.Fatal("call sum api %v", err)
	}

	log.Printf("result: %v", resp)
}