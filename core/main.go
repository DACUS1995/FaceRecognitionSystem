package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	facedetector "github.com/DACUS1995/FaceRecognition/core/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/core/sampler"
)

const (
	address   = "localhost:50051"
	imagePath = "./test_images/faces.jpg"
)

var close = make(chan bool)
var wg = sync.WaitGroup{}

var imageShape = []int32{349, 620, 3}

func main() {
	RunLocalImageFaceDetection(imagePath)

	go RunPeriodicDetection(5000, close)

	<-close
}

func gracefulExit() {
	close <- true
}

func RunPeriodicDetection(miliseconds int, close chan bool) {
	ticker := time.NewTicker(time.Duration(miliseconds) * time.Millisecond)
	client, err := facedetector.NewClient(address)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler := sampler.NewCameraSampler()

	for {
		select {
		case <-close:
			return
		case <-ticker.C:
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
	}
}

func RunLocalImageFaceDetection(imagePath string) {
	client, err := facedetector.NewClient(address)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler := sampler.NewLocalSampler(imagePath)
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
