package main

import (
	"fmt"
	"log"
	"time"

	facedetector "github.com/DACUS1995/FaceRecognition/coordinator/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/coordinator/sampler"
)

const (
	address   = "localhost:50051"
	imagePath = "./test_images/faces.jpg"
)

var quit = make(chan bool)

var imageShape = []int32{349, 620, 3}

func main() {
	RunLocalImageFaceDetection(imagePath)

	go RunPeriodicDetection(5000, quit)
}

func gracefulExit() {
	quit <- true
}

func RunPeriodicDetection(miliseconds int, quit chan bool) {
	ticker := time.NewTicker(time.Duration(miliseconds) * time.Millisecond)

	for {
		select {
		case <-quit:
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
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

	_, detectedFacesEmbeddings, err := client.DetectFaces(data, imageShape)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Number of faces detected: %v", len(detectedFacesEmbeddings)/125)
}
