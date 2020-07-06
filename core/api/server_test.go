package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGenerateEmbedding(t *testing.T) {
	trueEmbedding := []float32{0, 1, 2, 3, 4, 5}
	localImage, _ := os.Open("api\\test_images\\dr-house.jpeg")
	defer localImage.Close()

	ts := httptest.NewServer(http.HandlerFunc(generateEmbedding))
	defer ts.Close()

	client := ts.Client()
	client.Post(ts.URL, "image/jpeg", localImage)

	res, err := client.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	embedding, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(trueEmbedding, embedding) {
		t.Error("Embeddings are not equal")
	}
}
