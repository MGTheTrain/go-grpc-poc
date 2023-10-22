package main

import (
	"io"
	"log"
	"net"

	pb "upload_service"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedFileUploadServiceServer
}

func (s *server) UploadFile(stream pb.FileUploadService_UploadFileServer) error {
	fileData := []byte{}
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			log.Println("About to upload a file")
			// All chunks received
			return stream.SendAndClose(&pb.FileUploadResponse{Message: "File uploaded successfully"})
		}
		if err != nil {
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	maxMessageSize := 200 * 1024 * 1024 // 200 MB in bytes

	s := grpc.NewServer(
		// enable large file streaming uploads up to 200 MB in bytes
		grpc.MaxSendMsgSize(maxMessageSize),
		grpc.MaxRecvMsgSize(maxMessageSize),
	)
	pb.RegisterFileUploadServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
