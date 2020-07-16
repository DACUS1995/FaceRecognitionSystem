package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/DACUS1995/FaceRecognition/core/config"
	"github.com/DACUS1995/FaceRecognition/core/dbactions"
	facedetector "github.com/DACUS1995/FaceRecognition/core/face_detector"
)

// TODO Use a pool for the database connections

var Config *config.ConfigType = nil
var DatabaseClient *dbactions.DatabaseClient = nil
var FaceDetectorClient *facedetector.Client = nil

func StartServer(databaseClient dbactions.DatabaseClient) {
	Config = config.GetConfig()
	DatabaseClient = &databaseClient
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/embeddings", getAllEmbeddings)
	router.HandleFunc("/embedding/generate", generateEmbedding).Methods("POST")
	router.HandleFunc("/embedding", createNewEmbedding).Methods("POST")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func generateEmbedding(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	tee := io.TeeReader(r.Body, &buf)

	imageConfig, _, err := image.DecodeConfig(tee)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imageShape := []int32{int32(imageConfig.Width), int32(imageConfig.Height)}
	img, _ := ioutil.ReadAll(tee)

	detectedFacesEmbeddingsText, err := computeEmbedding(img, imageShape)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	detectedFacesEmbeddingsJoined := strings.Join(detectedFacesEmbeddingsText, ",")

	fmt.Fprintf(w, detectedFacesEmbeddingsJoined)
	fmt.Println("Generated a new embedding")
}

func createNewEmbedding(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	newRecord := dbactions.DatabaseRecord{}

	err := json.Unmarshal(reqBody, &newRecord)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Added a new embedding")
}

func getAllEmbeddings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllEmbeddings")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func computeEmbedding(img []byte, imageShape []int32) ([]string, error) {
	facedetectorClient := getFaceDetectorClient()

	_, detectedFacesEmbeddings, err := facedetectorClient.DetectFaces(img, imageShape)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.WithStack(err)
	}

	detectedFacesEmbeddingsText := []string{}
	for i := range detectedFacesEmbeddings {
		numberFloat := detectedFacesEmbeddings[i]
		textFloat := fmt.Sprintf("%f", numberFloat)
		detectedFacesEmbeddingsText = append(detectedFacesEmbeddingsText, textFloat)
	}
	return detectedFacesEmbeddingsText, nil
}

func getFaceDetectorClient() *facedetector.Client {
	if FaceDetectorClient == nil {
		faceDetectorClient, err := facedetector.NewClient(*Config.FaceDetectionServiceAddress)
		FaceDetectorClient = faceDetectorClient

		if err != nil {
			log.Panicf("Failed to instantiate client: %v", err)
		}
	}

	return FaceDetectorClient
}