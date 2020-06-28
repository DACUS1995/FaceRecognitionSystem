package dbactions

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBDatabaseClient struct {
	pathToDB      string
	levelDBClient *leveldb.DB
}

func AddRecord(name string, embedding []float32) {

}

func SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32) {
	return []DatabaseRecord{}, []float32{}
}

func GetAll() []DatabaseRecord {
	return []DatabaseRecord{}
}

func Save() {

}

func Load() {

}

func (dbClient *LevelDBDatabaseClient) init() {
	db, err := leveldb.OpenFile(dbClient.pathToDB, nil)
	if err != nil {
		log.Panicf("Failed to connect to db: %v", err)
	}

	dbClient.levelDBClient = db
}
