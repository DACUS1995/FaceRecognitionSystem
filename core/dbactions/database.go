package dbactions

type DatabaseClient interface {
	AddRecord(name string, embedding []float32) error
	SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32)
	GetAll() []DatabaseRecord
	Save()
	Load()
	Close()
}

type DatabaseRecord struct {
	Name      string
	Embedding []float32
}
