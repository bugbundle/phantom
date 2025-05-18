package camera

import "gocv.io/x/gocv"

type CameraService interface {
	Open(deviceId int) error
	Close() error
	CaptureImage() (*gocv.Mat, error)
}
