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

func (client *Client) DetectFaces(image []byte, imageShape []int32) ([]int32, []float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.grpcClient.DetectFaces(ctx, &Request{Image: image, ImageShape: imageShape})
	if err != nil {
		return nil, nil, err
	}

	return response.GetDetectedFacesBoxes(), response.GetDetectedFacesEmbeddings(), nil
}

func (client *Client) Close() {
	client.Close()
}
