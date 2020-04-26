package actions

type DatabaseClient interface {
	AddRecord()
	SearchRecordByName()
	SearchRecordBySimilarity()
}

type JSONDatabase struct {
	savePath        string
	numberOfRecords int
}

type JSONDatabaseClient struct {
	database JSONDatabase
}

func NewJSONDatabaseClient(path string) DatabaseClient {

}

func (client *JSONDatabaseClient) AddRecord()                {}
func (client *JSONDatabaseClient) SearchRecordByName()       {}
func (client *JSONDatabaseClient) SearchRecordBySimilarity() {}

func (database *JSONDatabase) loadDatabase() {

}

func (database *JSONDatabase) saveDatabase() {

}
