package main

import (
	"fmt"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	harrcascade := "opencv_haarcascade_frontalface_default.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	color := color.RGBA{0, 255, 0, 0}
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the device")
			continue
		}

		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			fmt.Println("detected", r)
			gocv.Rectangle(&img, r, color, 3)
		}

		window.IMShow(img)
		window.WaitKey(50)
	}
}
