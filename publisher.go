package main

// var publisherCmd = &cobra.Command{
// 	Use: "publisher",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		startParentChildProcess()
// 	},
// }

// func init() {
// 	RootCmd.AddCommand(publisherCmd)
// }

// var addr = fmt.Sprintf("127.0.0.1:%d", consumerPort)

// func startParentChildProcess() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	startSubProcess(ctx)

// 	conn := startConn()
// 	defer conn.Close()

// 	c := pb.NewMailAnalysisClient(conn)

// 	for keepLooping, timeout := true, time.After(10*time.Second); keepLooping; {
// 		select {
// 		case <-timeout:
// 			fmt.Println("timeout")
// 			cancel()

// 			keepLooping = false
// 		default:
// 			r, err := c.SimpleProcedure(ctx, &pb.SimpleRequest{Text: "ClientName"})
// 			if err != nil {
// 				log.Fatalf("could not greet: %v", err)
// 			}
// 			log.Printf("Get reply from consumer: %s", r.Text)

// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// 	fmt.Print("Finished")
// }

// func startSubProcess(ctx context.Context) {
// 	var out bytes.Buffer

// 	cmd := exec.CommandContext(ctx, "./main2", "consumer")

// 	cmd.Stderr = &out
// 	cmd.Stdout = &out

// 	// fmt.Println("Starting server..")

// 	if err := cmd.Start(); err != nil {
// 		log.Fatal(err)
// 	}

// 	go func() {
// 		reader := bufio.NewReader(&out)

// 		for keepGoing, timeout := true, ctx.Done(); keepGoing; {
// 			select {
// 			case <-timeout:
// 				keepGoing = false
// 			default:
// 				line, err := reader.ReadString('\n')
// 				if err == nil {
// 					fmt.Printf("server: %s", line)

// 				}

// 				time.Sleep(1 * time.Second)
// 			}
// 		}
// 	}()
// }

// func startConn() *grpc.ClientConn {
// 	conn, err := grpc.Dial(addr,
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithBlock(),
// 	)
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	return conn
// }
