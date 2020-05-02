package sampler

import (
	"io/ioutil"
	"log"
)

var imageShape = []int32{349, 620, 3}

type Sampler interface {
	Sample() ([]byte, []int32, error)
}

type localSampler struct {
	imagePath string
}

func NewLocalSampler(imagePath string) Sampler {
	return &localSampler{imagePath}
}

func (sampler *localSampler) Sample() ([]byte, []int32, error) {
	data, err := ioutil.ReadFile(sampler.imagePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", sampler.imagePath)
		return nil, nil, err
	}

	return data, imageShape, err
}
