package camera

import "gocv.io/x/gocv"

type CameraService interface {
	Open() error
	Close() error
	CaptureImage() (*gocv.Mat, error)
}
