package main

import (
	"fmt"

	pb "aflower/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := pb.NewSvrflowerSvrClient(conn)

	req := new(pb.HelloReq)
	req.Name = "kratos grpc"
	r, err := c.SayHelloURL(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Content)

}
