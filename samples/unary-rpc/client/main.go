package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "hello" // Import the generated package

	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	var name string

	var rootCmd = &cobra.Command{Use: "myclitool", Short: "CLI tool for sending an gRPC request"}
	rootCmd.Flags().StringVarP(&name, "name", "n", "Hello-Default-Value", "Some name")
	rootCmd.MarkFlagRequired("name")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Greeting)
}
