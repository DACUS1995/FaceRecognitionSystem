package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type DatabaseClient interface {
	AddRecord(name string, embedding []float32)
	SearchRecordBySimilarity(faceEmbedding []float32) []DatabaseRecord
	GetAll() []DatabaseRecord
	Save()
}

type JSONDatabase struct {
	savePath         string
	numberOfRecords  int
	recordCollection []DatabaseRecord
}

type DatabaseRecord struct {
	ID        int
	Name      string
	Embedding []float32
}

type JSONDatabaseClient struct {
	database *JSONDatabase
}

var databasePath string = "./"
var database *JSONDatabase = nil

func NewJSONDatabaseClient() DatabaseClient {
	if database == nil {
		database = &JSONDatabase{databasePath, 0, []DatabaseRecord{}}
	}

	return &JSONDatabaseClient{database}
}

func (client *JSONDatabaseClient) AddRecord(name string, embedding []float32) {
	client.database.recordCollection = append(
		client.database.recordCollection,
		DatabaseRecord{
			len(client.database.recordCollection),
			name,
			embedding,
		},
	)
}

func (client *JSONDatabaseClient) GetAll() []DatabaseRecord {
	return client.database.recordCollection
}

func (client *JSONDatabaseClient) Save() {
	client.database.saveDatabase()
}

func (client *JSONDatabaseClient) SearchRecordBySimilarity(faceEmbedding []float32) []DatabaseRecord {
	result := []DatabaseRecord{}

	for _, record := range client.database.recordCollection {
		distance, _ := cosineDistance(faceEmbedding, record.Embedding)

		if distance > 0.8 {
			result = append(result, record)
		}
	}

	return result
}

func (database *JSONDatabase) loadDatabase() {
	jsonFile, err := os.Open(database.savePath + "database.json")
	if err != nil {
		fmt.Printf("Failed to load database records from %v", database.savePath)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &database.recordCollection)
}

func (database *JSONDatabase) saveDatabase() {
	jsonContent, _ := json.Marshal(database.recordCollection)

	ioutil.WriteFile(database.savePath+"database.json", jsonContent, 0644)
}

func cosineDistance(vecA []float32, vecB []float32) (float32, error) {
	if len(vecA) != len(vecB) {
		return 0, errors.New("The vectors don't have the same length.")
	}

	var dotProductSum float32 = 0
	var magnitudeSumA float32 = 0
	var magnitudeSumB float32 = 0

	for idx := 0; idx < len(vecA); idx++ {
		dotProductSum += vecA[idx] * vecB[idx]
		magnitudeSumA += vecA[idx] * vecA[idx]
		magnitudeSumB += vecB[idx] * vecB[idx]
	}

	distance := float64(dotProductSum) / (math.Sqrt(float64(magnitudeSumA)) * math.Sqrt(float64(magnitudeSumB)))
	return float32(distance), nil
}
