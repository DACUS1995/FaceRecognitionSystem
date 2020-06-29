package dbactions

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBDatabaseClient struct {
	pathToDB string
	database *leveldb.DB
}

func (client *LevelDBDatabaseClient) AddRecord(name string, embedding []float32) error {
	byteEmbedding := embeddingToByte(embedding)

	err := client.database.Put([]byte(name), []byte(byteEmbedding), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to add new record.")
	}

	return nil
}

func (client *LevelDBDatabaseClient) SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32) {
	return []DatabaseRecord{}, []float32{}
}

func (client *LevelDBDatabaseClient) GetAll() []DatabaseRecord {
	return []DatabaseRecord{}
}

func (client *LevelDBDatabaseClient) Save() {
}

func (client *LevelDBDatabaseClient) Load() {
	client.init()
}

func (client *LevelDBDatabaseClient) Close() {
	client.database.Close()
}

func (client *LevelDBDatabaseClient) init() {
	db, err := leveldb.OpenFile(client.pathToDB, nil)
	if err != nil {
		log.Panicf("Failed to connect to db: %v", err)
	}

	client.database = db
}

func embeddingToByte(embedding []float32) []byte {
	buf := make([]byte, len(embedding)*4)

	for idx, val := range embedding {
		binary.BigEndian.PutUint32(buf[idx*4:(idx+1)*4], math.Float32bits(val))
	}

	return buf
}
