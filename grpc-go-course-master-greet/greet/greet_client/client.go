package main

import (
	"com.grpc.sabit/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	repeatFromServer(c)
}

func repeatFromServer(c greetpb.GreetServiceClient) {
	ctx := context.Background()
	fmt.Print("Ur number:")
	var num int64
	fmt.Scan(&num)
	req := &greetpb.DivideRequest{Greeting: &greetpb.Dividing{
		N: num,
	}}
	stream, err := c.Divide(ctx, req)
	if err != nil {
		log.Fatalf("error while calling Divide RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error while reciving from Divide RPC %v", err)
		}
		log.Printf("response from Divide:%v \n", res.GetResult())
	}

}