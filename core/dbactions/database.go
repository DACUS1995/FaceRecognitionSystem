package dbactions

type DatabaseClient interface {
	AddRecord(name string, embedding []float32)
	SearchRecordBySimilarity(faceEmbedding []float32) ([]DatabaseRecord, []float32)
	GetAll() []DatabaseRecord
	Save()
	Load()
}

type DatabaseRecord struct {
	ID        int
	Name      string
	Embedding []float32
}
