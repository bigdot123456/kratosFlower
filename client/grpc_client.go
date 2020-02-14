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

	r1,err1 := c.SayHello(context.Background(), req)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(r1)

	req2 := new(pb.Req)
	req2.Name="bigdot quickly"

	r2,err2 := c.Create(context.Background(), req2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(r2.Content)

}
