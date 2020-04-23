package face_detector

import (
	context "context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	address    string
	connection *grpc.ClientConn
	grpcClient FaceDetectorClient
}

func NewClient(address string) (*Client, error) {
	newGrpcConnection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return nil, err
	}

	grpcClient := NewFaceDetectorClient(newGrpcConnection)

	return &Client{
		address,
		newGrpcConnection,
		grpcClient,
	}, nil
}

func (client *Client) DetectFaces(image []byte, fileName string) ([]int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.grpcClient.DetectFaces(ctx, &Request{Image: image, Filename: fileName})
	if err != nil {
		return nil, err
	}

	return response.GetDetectedFaces(), err
}

func (client *Client) Close() {
	client.Close()
}
