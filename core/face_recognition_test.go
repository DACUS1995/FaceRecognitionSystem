package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/DACUS1995/FaceRecognition/core/dbactions"
	facedetector "github.com/DACUS1995/FaceRecognition/core/face_detector"
	sampler "github.com/DACUS1995/FaceRecognition/core/sampler"
)

func TestFaceRecognition(t *testing.T) {
	facedetectorClient, err := facedetector.NewClient(address)
	if err != nil {
		panic("Failed to instantiate facedetection client.")
	}
	sampler := sampler.NewLocalSampler(testImagePath)
	data, err := sampler.Sample()

	if err != nil {
		panic("Failed to sample the test image")
	}

	_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(data, imageShape)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	databaseClient := dbactions.NewJSONDatabaseClient()
	databaseClient.AddRecord("Anchorman", detectedFacesEmbeddings[:EMBEDDING_VECTOR_SIZE])
	databaseClient.Save()

	_, detectedFacesEmbeddings, err = facedetectorClient.DetectFaces(data, imageShape)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	found := false

	for idx := 0; idx < len(detectedFacesEmbeddings)-EMBEDDING_VECTOR_SIZE; idx += EMBEDDING_VECTOR_SIZE {
		if records, similarities := databaseClient.SearchRecordBySimilarity(detectedFacesEmbeddings[idx : idx+EMBEDDING_VECTOR_SIZE]); len(records) > 0 {
			fmt.Printf("Found: %v records\n", len(records))

			for i, record := range records {
				fmt.Printf("-> Record[%v]: %v | similarity: %v\n", i, record.Name, similarities[i])
			}
			found = true
		}
	}

	if found == false {
		t.Errorf("No matches found")
	}
}
