package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yoelsusanto/local-grpc/protobuf"
	"github.com/yoelsusanto/local-grpc/work"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnsafeMailAnalysisServer
}

func (s *Server) SimpleProcedure(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	return &pb.SimpleResponse{
		Text: work.SimpleProcedure(in.Text),
	}, nil
}

func (s *Server) HashProcedure(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	return &pb.SimpleResponse{
		Text: work.HashProcedure(in.Text),
	}, nil
}

func StartGRPCServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMailAnalysisServer(s, &Server{})

	log.Printf("listening on port: %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
