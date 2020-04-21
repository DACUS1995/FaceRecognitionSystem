package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/DACUS1995/FaceRecognition/Coordinator/face_detector"

	"google.golang.org/grpc"
)

const (
	address   = "localhost:50051"
	imagePath = "./test_images/faces.jpg"
)

func createFaceDetectorGRPCChannel() {

}

func testFaceDetector() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFaceDetectorClient(conn)

	// Contact the server and print out its response.
	path := imagePath
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", path)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	r, err := client.DetectFaces(ctx, &pb.Request{Image: data, Filename: "picture.jpg"})
	if err != nil {
		log.Fatalf("Could not Detect: %v", err)
	}
	log.Printf("Response: %v", r.GetDetectedFaces())
}

func main() {
	testFaceDetector()
}
