package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitServer()
}

func InitServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/embeddings", getAllEmbeddings)
	router.HandleFunc("/embedding", createNewEmbedding).Methods("POST")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func createNewEmbedding(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)

	fmt.Println("Added a new embedding")
}

func getAllEmbeddings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllEmbeddings")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
