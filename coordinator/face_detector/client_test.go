package face_detector

import (
	"io/ioutil"
	"log"
	"testing"
)

const (
	address   = "localhost:50051"
	imagePath = "../test_images/faces.jpg"
)

func TestFaceDetector(t *testing.T) {
	client, err := NewClient(address)
	if err != nil {
		t.Error("Failed to instantiate the face detector client.")
	}

	// Contact the server and print out its response.
	path := imagePath

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", path)
	}

	expected := []int32{381, 50, 424, 94, 118, 67, 170, 118, 469, 84, 521, 136, 245, 32, 296, 84}
	response, err := client.DetectFaces(data, "picture.jpg")

	if len(response) != len(expected) {
		t.Errorf("Expected: [%v] | Returned: [%v]", expected, response)
	}
}
