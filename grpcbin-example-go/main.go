package main

import (
	"fmt"

	pb "github.com/moul/pb/grpcbin/go-grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)
	fmt.Println(client)
}
