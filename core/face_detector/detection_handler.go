package face_detector

import (
	"log"

	"github.com/DACUS1995/FaceRecognition/core/config"
	"github.com/DACUS1995/FaceRecognition/core/dbactions"
)

type Handler interface {
	Handle([]int32, []float32)
}

type HandlerParameter struct {
	boundingBoxes []int32
	embeddings    []float32
}

type DatabaseSearcher struct {
	databaseClient dbactions.DatabaseClient
}

func NewDatatabaseSeacher(databaseClient dbactions.DatabaseClient) Handler {
	return &DatabaseSearcher{databaseClient}
}

func (handler *DatabaseSearcher) Handle(boundingBoxes []int32, detectedFacesEmbeddings []float32) {
	Config := config.GetConfig()

	for idx := 0; idx < len(detectedFacesEmbeddings)-*Config.EmbeddingVectorSize; idx += *Config.EmbeddingVectorSize {
		if records, similarities := handler.databaseClient.SearchRecordBySimilarity(detectedFacesEmbeddings[idx : idx+*Config.EmbeddingVectorSize]); len(records) > 0 {
			for i, record := range records {
				log.Printf("-> Record[%v] | similarity: %v\n", record.Name, similarities[i])
			}
		}
	}
}
