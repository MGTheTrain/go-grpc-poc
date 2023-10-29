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
	address     = "localhost:8080"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	name := defaultName

	var rootCmd = &cobra.Command{Use: "myclitool", Short: "CLI tool for sending an gRPC request"}
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", defaultName, "Some name")
	rootCmd.MarkPersistentFlagRequired("name")
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
