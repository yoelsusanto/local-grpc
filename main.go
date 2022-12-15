package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yoelsusanto/local-grpc/grpc"
)

var RootCmd = &cobra.Command{
	Use: "main",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main")
	},
}

func main() {
	RootCmd.Execute()
}

var grpcCmd = &cobra.Command{
	Use:  "grpc",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal("Failed to parse port", err)
		}

		grpc.StartGRPCServer(int(port))
	},
}

func init() {
	RootCmd.AddCommand(grpcCmd)
}
