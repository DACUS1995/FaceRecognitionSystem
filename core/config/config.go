package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigType struct {
	FaceDetectionServiceAddress *string `json:"face-detection-service-address"`
	CameraSamplerServiceAddress *string `json:"camera-sampler-service-address"`
	EmbeddingVectorSize         *int    `json:"embedding-vector-size"`
	SamplingIntervalMiliseconds *int    `json:"sampling-interval-miliseconds"`
	TestImagePath               *string `json:"test-image-path"`
}

var config *ConfigType = nil

func GetConfig() *ConfigType {
	if config == nil {
		loadConfig()
	}

	return config
}

func loadConfig() {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}
	defer jsonFile.Close()

	byteValues, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValues, &config)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Loaded config file.")
}
