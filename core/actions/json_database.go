package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// TODO make database a singleton

type DatabaseClient interface {
	AddRecord(name string, embedding []float32)
	SearchRecordByName()
	SearchRecordBySimilarity()
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

func NewJSONDatabaseClient(path string) DatabaseClient {
	return &JSONDatabaseClient{
		&JSONDatabase{path, 0, []DatabaseRecord{}},
	}
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

func (client *JSONDatabaseClient) SearchRecordByName()       {}
func (client *JSONDatabaseClient) SearchRecordBySimilarity() {}

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
