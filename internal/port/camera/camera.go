package camera

import (
	"fmt"
	"sync"

	"github.com/bugbundle/phantom/internal/adapter/camera"
	"gocv.io/x/gocv"
)

type webCamSingleton struct {
	capture *gocv.VideoCapture
	isOpen  bool
	mu      sync.Mutex
}

var (
	instance *webCamSingleton
	once     sync.Once
)

// CaptureImage captures an image from the webcam and returns it.
func (w *webCamSingleton) CaptureImage() (*gocv.Mat, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if !w.isOpen {
		return nil, fmt.Errorf("webcam unavailable")
	}
	img := gocv.NewMat()
	if ok := w.capture.Read(&img); !ok {
		return nil, fmt.Errorf("failed to capture image")
	}
	return &img, nil
}

func GetInstance() camera.CameraService {
	once.Do(func() {
		instance = &webCamSingleton{}
	})
	return instance
}

// Open opens the webcam if not already open.
func (w *webCamSingleton) Open() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.isOpen {
		return fmt.Errorf("webcam already open")
	}
	capture, err := gocv.OpenVideoCapture(0)
	if err != nil {
		return fmt.Errorf("error opening webcam: %v", err)
	}
	w.capture = capture
	w.isOpen = true
	return nil
}

// Close closes the webcam and releases the resource.
func (w *webCamSingleton) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.isOpen && w.capture != nil {
		w.capture.Close()
		w.isOpen = false
		w.capture = nil
	}
	// TODO: Add error management later
	return nil
}
