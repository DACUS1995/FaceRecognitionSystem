package dbactions

import (
	"testing"
)

func TestJSONDatabase(t *testing.T) {
	databaseClient := NewJSONDatabaseClient()
	databaseClient.AddRecord("Tudor", []float32{0.1, 0.2})

	databaseClient.Save()
	collection := databaseClient.GetAll()

	if len(collection) != 1 {
		t.Errorf("Expected: [%v] | Returned: [%v]", 1, len(collection))
	}
}
