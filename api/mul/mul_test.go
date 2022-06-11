package mul

import (
	"context"
	"testing"

	_ "github.com/kanguki/grpc-microservices-example/api/div" //test if proto register conflicts
	"github.com/kanguki/grpc-microservices-example/api/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func _TestMul(t *testing.T) {
	TARGET := "localhost:4002"
	conn, err := grpc.Dial(TARGET, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Log("failed to dial server: %v", err)
	}
	client := NewMulClient(conn)
	res, err := client.Do(context.Background(), &Request{Term1: 1, Term2: 2})
	if err != nil {
		//this requires external network calls so I ignore 😛
		return
	}
	t.Log(*res)
}
