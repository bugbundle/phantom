package utils

import (
	"errors"
	"fmt"

	"gocv.io/x/gocv"
)

var instance *WebCamSingleton

type WebCamSingleton struct {
	capture *gocv.VideoCapture
	isOpen  bool
}

func GetCamera() (*WebCamSingleton, error) {
	if instance != nil {
		return instance, nil
	}
	return nil, errors.New("camera is missing")
}

func CreateOrGetCamera() *WebCamSingleton {
	if instance == nil {
		instance = &WebCamSingleton{}
	}
	return instance
}

func DeleteCamera() error {
	if instance == nil {
		return nil
	}

	instance.Stop()
	instance = nil

	return nil
}

func (wc *WebCamSingleton) IsOpen() bool {
	return wc.isOpen
}

// Open starts the webcam capture.
func (wc *WebCamSingleton) Open() error {
	if wc.isOpen {
		return fmt.Errorf("webcam is already open")
	}

	capture, err := gocv.OpenVideoCapture(0) // Use 0 for default webcam
	if err != nil {
		return fmt.Errorf("error opening video capture device: %v", err)
	}

	wc.capture = capture
	wc.isOpen = true
	return nil
}

// Stop stops the webcam capture.
func (wc *WebCamSingleton) Stop() {
	if !wc.isOpen {
		fmt.Println("webcam is not open")
		return
	}

	wc.capture.Close()
	wc.isOpen = false
}

// CaptureImage captures an image from the webcam and returns it.
func (wc *WebCamSingleton) CaptureImage() (*gocv.Mat, error) {
	if !wc.isOpen {
		return nil, fmt.Errorf("webcam is not open")
	}

	img := gocv.NewMat()
	if ok := wc.capture.Read(&img); !ok {
		return nil, fmt.Errorf("failed to read image from webcam")
	}
	return &img, nil
}
