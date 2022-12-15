package main

// var hasher = sha256.New()

// var consumerCmd = &cobra.Command{
// 	Use: "consumer",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		StartConsumer()
// 	},
// }

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// 	RootCmd.AddCommand(consumerCmd)
// }

// type Server struct {
// 	pb.UnsafeMailAnalysisServer
// }

// func (s *Server) SimpleProcedure(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
// 	return &pb.SimpleResponse{
// 		Text: in.Text,
// 	}, nil
// }

// func SimpleProcedure(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
// 	return &pb.SimpleResponse{
// 		Text: in.Text,
// 	}, nil
// }

// func (s *Server) HashProcedure(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
// 	hasher.Reset()
// 	for i := 0; i < 100000; i++ {
// 		hasher.Write([]byte(in.Text))
// 	}
// 	hasher.Sum(nil)

// 	return &pb.HashResponse{
// 		Text: in.Text,
// 	}, nil
// }

// func HashProcedure(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
// 	hasher.Reset()
// 	for i := 0; i < 100000; i++ {
// 		hasher.Write([]byte(in.Text))
// 	}
// 	hasher.Sum(nil)

// 	return &pb.HashResponse{
// 		Text: in.Text,
// 	}, nil
// }

// const consumerPort = 50001

// func StartConsumer() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", consumerPort))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterMailAnalysisServer(s, &Server{})

// 	// log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
