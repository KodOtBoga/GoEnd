package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"com.grpc.sabit/greet/greetpb"
	"google.golang.org/grpc"
)

//Server with embedded UnimplementedDivideServiceServer
type Server struct {
	greetpb.UnimplementedDivideServiceServer
}


func (s *Server) DivideManyTimes(req *greetpb.DivideRequest, stream greetpb.DivideService_DivideServer) error {

	number := int(req.GetDividing().GetNumber())
	for number > 1 {
		for i := 2; number >= i;{
			if number % i == 0{
				number = number / i
				res := &greetpb.DivideResponse{Result: fmt.Sprintf(strconv.Itoa(i))}
				if err := stream.Send(res); err != nil {
					log.Fatalf("error while sending greet many times responses: %v", err.Error())
				}
				time.Sleep(time.Second)
				i = number
			}
			i = i + 1
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
