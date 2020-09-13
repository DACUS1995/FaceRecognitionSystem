package dbactions

import "errors"

type DatabaseClient interface {
	AddRecord(name string, embedding []float32) error
	SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32, error)
	GetAll() ([]DatabaseRecord, error)
	Save() error
	Load() error
	Close() error
}

type DatabaseRecord struct {
	Name      string    `json:"name"`
	Embedding []float32 `json:"embedding"`
}

var ErrDatabaseClosed = errors.New("Database is closed")
var ErrNoMatchesFound = errors.New("Cannot find any matches for the provided embedding")
