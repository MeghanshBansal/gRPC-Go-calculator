package main

import (
	pb "MeghanshBansal/calculator/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.NewClient("localhost:50012", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("failed to dial connection to server")
	}
	defer conn.Close()
	
	client := pb.NewCalculatorClient(conn)
	go callAdd(&client)
	go callSubTract(&client)
	go callMultiply(&client)
	go callDivision(&client)
	time.Sleep(1  * time.Second)
}


func callAdd(client *pb.CalculatorClient) {
	res, err := (*client).Add(context.Background(), &pb.CalculatorRequest{Num1: 1, Num2: 2})
	if err != nil{
		log.Println("failed the add api with err", err)
		return 
	}
	fmt.Println("Addition is ", res.Res)
}

func callSubTract(client *pb.CalculatorClient){
	res, err := (*client).Subtract(context.Background(), &pb.CalculatorRequest{Num1: 1, Num2: 2})
	if err != nil{
		log.Println("failed the sub api with err", err)
		return 
	}
	fmt.Println("subtraction is ", res.Res)
}

func callMultiply(client *pb.CalculatorClient) {
	res, err := (*client).Multiply(context.Background(), &pb.CalculatorRequest{Num1: 1, Num2: 2})
	if err != nil{
		log.Println("failed the multiplication api with err", err)
		return 
	}
	fmt.Println("multiplication is ", res.Res)
}

func callDivision(client *pb.CalculatorClient){
	res, err := (*client).Divide(context.Background(), &pb.CalculatorRequest{Num1: 1, Num2: 2})
	if err != nil{
		log.Println("failed the div api with err", err)
		return 
	}
	fmt.Println("division is ", res.Res)

	res, err = (*client).Divide(context.Background(), &pb.CalculatorRequest{Num1: 1, Num2: 0})
	if err != nil{
		log.Println("failed the div api with err", err)
		return 
	}
	fmt.Println("division is ", res.Res)

}
