package main

import (
	"fmt"

	facedetector "github.com/DACUS1995/FaceRecognition/coordinator/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/coordinator/sampler"
)

const (
	address   = "localhost:50051"
	imagePath = "./test_images/faces.jpg"
)

func main() {
	RunLocalImageFaceDetection(imagePath)
	RunPeriodicDetection()
}

func RunPeriodicDetection() {

}

func RunLocalImageFaceDetection(imagePath string) {
	client, err := facedetector.NewClient(address)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler := sampler.NewOneTimeSampler(imagePath)
	data, err := sampler.Sample()

	if err != nil {
		panic("Failed to sample the test image")
	}

	response, err := client.DetectFaces(data, "picture.jpg")
	fmt.Printf("Number of faces detected: %v", len(response)/4)
}
