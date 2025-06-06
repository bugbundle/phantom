package controls

import (
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"time"

	"github.com/bugbundle/phantom/internal/port/camera"
	"gocv.io/x/gocv"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /cameras/{deviceId}", startCameraHandler)
	router.HandleFunc("DELETE /cameras", stopCameraHandler)
	router.HandleFunc("GET /cameras", StreamVideoHandler)
}

// If not yet, create a camera entity
func startCameraHandler(w http.ResponseWriter, r *http.Request) {
	g, err := strconv.Atoi(r.PathValue("deviceId"))
	if err != nil {
		http.Error(w, "Invalid deviceId", http.StatusNotFound)
	}
	// Trigger singleton to instanciate camera
	device := camera.GetInstance()
	device.Open(g)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// If existing, remove a camera entity
func stopCameraHandler(w http.ResponseWriter, r *http.Request) {
	device := camera.GetInstance()
	device.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// This function retrieve camera device and start streaming using multipart/x-mixed-replace
// TODO: Add device number option
func StreamVideoHandler(w http.ResponseWriter, r *http.Request) {
	// If the camera is unavailable return 428
	webcam := camera.GetInstance()

	mimeWriter := multipart.NewWriter(w)
	w.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", mimeWriter.Boundary()))
	partHeader := make(textproto.MIMEHeader)
	partHeader.Add("Content-Type", "image/jpeg")

	for {
		img, err := webcam.CaptureImage()
		if err != nil {
			http.Error(w, "Error capturing image: %v", http.StatusServiceUnavailable)
		}
		// defer img.Close()

		resizedFrame := gocv.NewMat()

		gocv.Resize(*img, &resizedFrame, image.Point{X: 256, Y: 144}, 0, 0, gocv.InterpolationDefault)

		buf, err := gocv.IMEncode(gocv.JPEGFileExt, resizedFrame)
		if err != nil {
			log.Println("Error encoding frame: ", err)
			continue
		}

		partWriter, _ := mimeWriter.CreatePart(partHeader)
		if _, err := partWriter.Write(buf.GetBytes()); err != nil {
			log.Println("Error while processing buffer")
		}

		// we want to record around 10 fps
		// this mean every second send 10 images
		// Let's assume reading, encoding and  writing do not consume any resources
		// If we sleep for 1/10 of a second we roughly approximate 10 fps
		time.Sleep(200 * time.Millisecond)
	}
}
