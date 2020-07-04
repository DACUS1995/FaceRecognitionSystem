package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DACUS1995/FaceRecognition/core/dbactions"
	"github.com/gorilla/mux"
)

func main() {
	InitServer()
}

func InitServer() {
	// TODO get a database injected
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/embeddings", getAllEmbeddings)
	router.HandleFunc("/embedding/generate", generateEmbedding).Methods("POST")
	router.HandleFunc("/embedding", createNewEmbedding).Methods("POST")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func generateEmbedding(w http.ResponseWriter, r *http.Request) {
	image, _ := ioutil.ReadAll(r.Body)
	// TODO: get image shape

	facedetectorClient, err := facedetector.NewClient(*Config.FaceDetectionServiceAddress)
	if err != nil {
		log.Panicf("Failed to instantiate client: %v", err)
	}

	_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(image, imageShape)
	if err != nil {
		log.Panicf("Error: %v", err)
	}

	fmt.Fprintf(w, "123123")
	fmt.Println("Generated a new embedding")
}

func createNewEmbedding(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	newRecord := dbactions.DatabaseRecord{}
	json.Unmarshal(reqBody, &newRecord)

	fmt.Println("Added a new embedding")
}

func getAllEmbeddings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllEmbeddings")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
