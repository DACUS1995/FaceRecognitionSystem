package sampler

import (
	"io/ioutil"
	"log"
)

type Sampler interface {
	Sample() ([]byte, error)
}

type localSampler struct {
	imagePath string
}

func NewLocalSampler(imagePath string) Sampler {
	return &localSampler{imagePath}
}

func (sampler *localSampler) Sample() ([]byte, error) {
	data, err := ioutil.ReadFile(sampler.imagePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", sampler.imagePath)
		return nil, err
	}

	return data, err
}

type cameraSampler struct {
}

func NewCameraSampler() Sampler {
	return &cameraSampler{}
}

func (sampler *cameraSampler) Sample() ([]byte, error) {
	panic("Must be implemented")
	// return nil, nil
}
