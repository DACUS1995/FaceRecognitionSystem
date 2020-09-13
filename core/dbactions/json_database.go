package dbactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sync"
)

type JSONDatabase struct {
	savePath         string
	numberOfRecords  int
	recordCollection []DatabaseRecord
	open             bool
}

type JSONDatabaseClient struct {
	database *JSONDatabase
	mux      sync.RWMutex
}

var databasePath string = "./"
var database *JSONDatabase = nil

func NewJSONDatabaseClient() DatabaseClient {
	if database == nil {
		database = &JSONDatabase{
			databasePath,
			0,
			[]DatabaseRecord{},
			true,
		}
	}

	return &JSONDatabaseClient{
		database: database,
	}
}

func (client *JSONDatabaseClient) AddRecord(name string, embedding []float32) error {
	client.mux.Lock()
	defer client.mux.Unlock()

	if client.isClosed() {
		return ErrDatabaseClosed
	}

	// TODO Store the embeddings indexed by name
	for _, record := range client.database.recordCollection {
		if record.Name == name {
			record.Embedding = embedding
			return nil
		}
	}

	client.database.recordCollection = append(
		client.database.recordCollection,
		DatabaseRecord{
			name,
			embedding,
		},
	)

	return nil
}

func (client *JSONDatabaseClient) GetAll() ([]DatabaseRecord, error) {
	client.mux.RLock()
	defer client.mux.RUnlock()

	if client.isClosed() {
		return nil, ErrDatabaseClosed
	}

	newRecordCollection := []DatabaseRecord{}
	copy(client.database.recordCollection, newRecordCollection)

	return newRecordCollection, nil
}

func (client *JSONDatabaseClient) Save() error {
	client.mux.RLock()
	defer client.mux.RUnlock()

	if client.isClosed() {
		return ErrDatabaseClosed
	}

	client.database.saveDatabase()
	return nil
}

func (client *JSONDatabaseClient) Load() error {
	client.mux.Lock()
	defer client.mux.Unlock()

	if client.isClosed() {
		return ErrDatabaseClosed
	}

	client.database.loadDatabase()
	return nil
}

func (client *JSONDatabaseClient) Close() error {
	client.mux.Lock()
	defer client.mux.Unlock()

	if client.isClosed() {
		return ErrDatabaseClosed
	}

	client.Save()
	client.database.open = false
	return nil
}

func (client *JSONDatabaseClient) SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32, error) {
	client.mux.RLock()
	defer client.mux.RUnlock()

	if client.isClosed() {
		return nil, nil, ErrDatabaseClosed
	}

	result := []DatabaseRecord{}
	similarities := []float32{}

	for _, record := range client.database.recordCollection {
		distance, _ := cosineDistance(faceEmbedding, record.Embedding)

		if distance > 0.8 {
			result = append(result, record)
			similarities = append(similarities, distance)
		}
	}

	if len(result) == 0 {
		return nil, nil, ErrNoMatchesFound
	}

	return result, similarities, nil
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

func (client *JSONDatabaseClient) isClosed() bool {
	return !client.database.open
}
