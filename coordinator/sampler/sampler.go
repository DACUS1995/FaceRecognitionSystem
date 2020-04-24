package sampler

import (
	"io/ioutil"
	"log"
)

type Sampler interface {
	Sample() ([]byte, error)
}

type oneTimeSampler struct {
	imagePath string
}

func NewOneTimeSampler(imagePath string) Sampler {
	return &oneTimeSampler{imagePath}
}

func (sampler *oneTimeSampler) Sample() ([]byte, error) {
	data, err := ioutil.ReadFile(sampler.imagePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", sampler.imagePath)
		return nil, err
	}

	return data, err
}
