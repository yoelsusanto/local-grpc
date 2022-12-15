package grpc

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"

	pb "github.com/yoelsusanto/local-grpc/protobuf"
	"github.com/yoelsusanto/local-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var resultSimple *pb.SimpleResponse

func BenchmarkSimpleGRPC(b *testing.B) {
	b.StopTimer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := 50000 + rand.Int31n(100)

	utils.StartSubProcess(ctx, "main2", "grpc", fmt.Sprintf("%d", port))

	conn := startConn(int(port))
	defer conn.Close()

	c := pb.NewMailAnalysisClient(conn)

	var (
		r   *pb.SimpleResponse
		err error
	)

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		r, err = c.SimpleProcedure(ctx, &pb.SimpleRequest{
			Text: "hello_world",
		})
		if err != nil {
			panic(err)
		}
	}

	b.StopTimer()

	resultSimple = r
}

func BenchmarkComplexGRPC(b *testing.B) {
	b.StopTimer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := 50000 + rand.Int31n(100)

	utils.StartSubProcess(ctx, "main2", "grpc", fmt.Sprintf("%d", port))

	conn := startConn(int(port))
	defer conn.Close()

	c := pb.NewMailAnalysisClient(conn)

	var (
		r   *pb.SimpleResponse
		err error
	)

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		r, err = c.HashProcedure(ctx, &pb.SimpleRequest{
			Text: "hello_world",
		})
		if err != nil {
			panic(err)
		}
	}

	b.StopTimer()

	resultSimple = r
}

func startConn(port int) *grpc.ClientConn {
	addr := fmt.Sprintf("127.0.0.1:%d", port)

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}
