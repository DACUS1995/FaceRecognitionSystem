package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	facedetector "github.com/DACUS1995/FaceRecognition/core/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/core/sampler"
)

type configType struct {
	FaceDetectionServiceAddress string `json:"face-detection-service-address"`
	CameraSamplerServiceAddress string `json:"camera-sampler-service-address"`
	EmbeddingVectorSize         int    `json:"embedding-vector-size"`
}

var config *configType = nil

const (
	testImagePath = "./test_images/faces.jpg"
)

var close = make(chan bool)
var wg = sync.WaitGroup{}

func main() {
	loadConfig()

	RunLocalImageFaceDetection(testImagePath)
	go RunPeriodicDetection(5000, close)

	<-close
}

func gracefulExit() {
	close <- true
}

func RunPeriodicDetection(miliseconds int, close chan bool) {
	ticker := time.NewTicker(time.Duration(miliseconds) * time.Millisecond)
	facedetectorClient, err := facedetector.NewClient(config.FaceDetectionServiceAddress)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler, err := sampler.NewCameraSampler(config.CameraSamplerServiceAddress)
	if err != nil {
		panic("Failed to create connection to the sampler.")
	}

	for {
		select {
		case <-close:
			return
		case <-ticker.C:
			data, imageShape, err := sampler.Sample()
			if err != nil {
				panic(err)
			}

			_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(data, imageShape)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			fmt.Printf("Number of faces detected: %v", len(detectedFacesEmbeddings)/config.EmbeddingVectorSize)
		}
	}
}

func RunLocalImageFaceDetection(testImagePath string) {
	facedetectorClient, err := facedetector.NewClient(config.FaceDetectionServiceAddress)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler := sampler.NewLocalSampler(testImagePath)
	data, imageShape, err := sampler.Sample()

	if err != nil {
		panic("Failed to sample the test image")
	}

	_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(data, imageShape)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Number of faces detected: %v", len(detectedFacesEmbeddings)/config.EmbeddingVectorSize)
}

func loadConfig() {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		panic("Failed to load config")
	}
	defer jsonFile.Close()

	byteValues, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValues, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded config file.")
}
