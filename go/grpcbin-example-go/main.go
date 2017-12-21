package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/moul/pb/grpcbin/go-grpc"
)

func main() {
	// dial
	// conn, _ := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	conn, _ := grpc.Dial("grpcb.in:9001", grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: "", InsecureSkipVerify: true})))
	defer conn.Close()

	// create client and context
	client := pb.NewGRPCBinClient(conn)
	ctx := context.Background()

	// call DummyUnary
	res, err := client.DummyUnary(ctx, &pb.DummyMessage{
		FString: "hello",
		FInt32:  42,
	})
	if err != nil {
		log.Fatalf("failed to call DummyUnary: %v", err)
	}
	fmt.Println(res)
}
