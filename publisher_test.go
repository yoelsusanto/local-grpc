package main

// var result *pb.SimpleResponse

// func BenchmarkSimpleRPC(b *testing.B) {
// 	b.StopTimer()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	startSubProcess(ctx)

// 	conn := startConn()
// 	defer conn.Close()

// 	c := pb.NewMailAnalysisClient(conn)

// 	var (
// 		r   *pb.SimpleResponse
// 		err error
// 	)

// 	b.StartTimer()

// 	for n := 0; n < b.N; n++ {
// 		r, err = c.SimpleProcedure(ctx, &pb.SimpleRequest{
// 			Text: "hello_world",
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	result = r
// }

// func BenchmarkSimpleDirect(b *testing.B) {
// 	b.StopTimer()

// 	var (
// 		r   *pb.SimpleResponse
// 		err error
// 	)

// 	b.StartTimer()

// 	for n := 0; n < b.N; n++ {
// 		r, err = SimpleProcedure(context.Background(), &pb.SimpleRequest{
// 			Text: "hello_world",
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	result = r
// }

// var resultHash *pb.HashResponse

// func BenchmarkHashRPC(b *testing.B) {
// 	b.StopTimer()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	startSubProcess(ctx)

// 	conn := startConn()
// 	defer conn.Close()

// 	c := pb.NewMailAnalysisClient(conn)

// 	var (
// 		r   *pb.HashResponse
// 		err error
// 	)

// 	b.StartTimer()

// 	for n := 0; n < b.N; n++ {
// 		r, err = c.HashProcedure(ctx, &pb.HashRequest{
// 			Text: "hello_world",
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	resultHash = r
// }

// func BenchmarkHashDirect(b *testing.B) {
// 	b.StopTimer()

// 	var (
// 		r   *pb.HashResponse
// 		err error
// 	)

// 	b.StartTimer()

// 	for n := 0; n < b.N; n++ {
// 		r, err = HashProcedure(context.Background(), &pb.HashRequest{
// 			Text: "hello_world",
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	resultHash = r
// }
