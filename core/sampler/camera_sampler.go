package sampler

import (
	context "context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cameraSampler struct {
	address    string
	connection *grpc.ClientConn
	grpcClient CameraSamplerClient
}

func NewCameraSampler(address string) (*cameraSampler, error) {
	newGrpcConnection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return nil, err
	}

	grpcClient := NewCameraSamplerClient(newGrpcConnection)

	return &cameraSampler{
		address,
		newGrpcConnection,
		grpcClient,
	}, nil
}

func (sampler *cameraSampler) Sample() ([]byte, []int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := sampler.grpcClient.SampleImage(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, nil, err
	}

	return response.GetImage(), response.GetImageShape(), nil
}
