package main

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "upload_service"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultPath = "D:/hdr-environments/BlueBackground.jpg"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileUploadServiceClient(conn)

	var path string

	var rootCmd = &cobra.Command{Use: "myclitool", Short: "CLI tool for sending an gRPC request"}
	rootCmd.Flags().StringVarP(&path, "path", "n", defaultPath, "Some path")
	rootCmd.MarkFlagRequired("path")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	stream, err := c.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("Could not upload file: %v", err)
	}
	chunk := &pb.FileChunk{Data: fileData}
	if err := stream.Send(chunk); err != nil {
		log.Fatalf("Error sending chunk: %v", err)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Server Response: %s", response.Message)
}
