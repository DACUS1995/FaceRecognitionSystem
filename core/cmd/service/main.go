package main

import (
	"log"
	"sync"
	"time"

	"github.com/DACUS1995/FaceRecognition/core/api"
	"github.com/DACUS1995/FaceRecognition/core/config"
	"github.com/DACUS1995/FaceRecognition/core/dbactions"
	facedetector "github.com/DACUS1995/FaceRecognition/core/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/core/sampler"
)

var close = make(chan bool)
var wg = sync.WaitGroup{}

func main() {
	config := config.GetConfig()
	databaseClient := GetDatabase()

	api.StartServer(databaseClient)

	if config.TestImagePath != nil {
		RunLocalImageFaceDetection(*config.TestImagePath)
	}

	go RunPeriodicDetection(
		*config.SamplingIntervalMiliseconds,
		close,
		databaseClient,
	)

	<-close
}

func gracefulExit() {
	log.Printf("Trying to gracefully exit.")
	close <- true
}

func GetDatabase() dbactions.DatabaseClient {
	databaseClient := dbactions.NewJSONDatabaseClient()
	databaseClient.Load()
	return databaseClient
}

func RunPeriodicDetection(miliseconds int, close chan bool, databaseClient dbactions.DatabaseClient) {
	ticker := time.NewTicker(time.Duration(miliseconds) * time.Millisecond)

	config := config.GetConfig()
	facedetectorClient, err := facedetector.NewClient(*config.FaceDetectionServiceAddress)
	if err != nil {
		log.Panicf("Failed to instantiate client: %v", err)
	}

	sampler, err := sampler.NewCameraSampler(*config.CameraSamplerServiceAddress)
	if err != nil {
		log.Panic("Failed to create connection to the sampler.")
	}

	detectionHandler := facedetector.NewDatatabaseSeacher(databaseClient)

	for {
		select {
		case <-close:
			return
		case <-ticker.C:
			data, imageShape, err := sampler.Sample()
			if err != nil {
				log.Println(err)
			}

			boundingBoxes, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(data, imageShape)
			if err != nil {
				log.Printf("Error: %v", err)
			}

			detectionHandler.Handle(boundingBoxes, detectedFacesEmbeddings)

			// log.Printf("Number of faces detected: %v", len(detectedFacesEmbeddings)/(*Config.EmbeddingVectorSize))
		}
	}
}

func RunLocalImageFaceDetection(testImagePath string) {
	config := config.GetConfig()
	facedetectorClient, err := facedetector.NewClient(*config.FaceDetectionServiceAddress)
	if err != nil {
		log.Panicf("Failed to instantiate client: %v", err)
	}
	sampler := sampler.NewLocalSampler(testImagePath)
	data, imageShape, err := sampler.Sample()

	if err != nil {
		log.Panicf("Failed to sample the test image: %v", err)
	}

	_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(data, imageShape)
	if err != nil {
		log.Panicf("Error: %v", err)
	}

	log.Printf("Number of faces detected: %v", len(detectedFacesEmbeddings)/(*config.EmbeddingVectorSize))

}
