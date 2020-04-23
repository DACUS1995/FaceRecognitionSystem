package main

import (
	ps "periodicsampler"
)

func main() {
	sampler = ps.New()
	client, err := NewClient(address)
	if err != nil {
		panic("Failed to instantiate client.")
	}
	sampler.AddHandler(func(image) {
		response, err := client.DetectFaces(data, "picture.jpg")
	})
}
