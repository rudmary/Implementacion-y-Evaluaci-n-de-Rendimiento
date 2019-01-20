
package main

import (
	"log"
	//"os"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

    start := time.Now()
    name := defaultName
	message, err := c.Ping(ctx,  &pb.PingRequest{Message: name})
	end := time.Now()
	fmt.Println(end.Sub(start))
	log.Println(message)
	
}