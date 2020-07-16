package main

import (
	"github.com/DACUS1995/FaceRecognition/core/api"
	"github.com/DACUS1995/FaceRecognition/core/dbactions"
)

func main() {
	databaseClient := dbactions.NewJSONDatabaseClient()
	databaseClient.Load()
	api.StartServer(databaseClient)
}
