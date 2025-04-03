package routes

import (
	"fmt"
	"html/template"
	"image"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/bugbundle/phantom/api/utils"
	"github.com/charmbracelet/log"
	"gocv.io/x/gocv"
)

// Return default Homepage, a simple alpineJS application to users
func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	template, err := template.ParseFiles("api/templates/index.html.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// This function retrieve camera device and start streaming using multipart/x-mixed-replace
// TODO: Add device number option
func StreamVideo(w http.ResponseWriter, r *http.Request) {
	// If the camera is unavailable return 428
	webcam, err := utils.GetCamera()
	if err != nil {
		http.Error(w, "Camera unavailable.", http.StatusPreconditionRequired)
		return
	}

	// Try to read the webcam
	openErr := webcam.Open()
	if openErr != nil {
		log.Debug(fmt.Sprintf("Got the following error: %d", openErr), http.StatusServiceUnavailable)
	}

	mimeWriter := multipart.NewWriter(w)
	w.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", mimeWriter.Boundary()))
	partHeader := make(textproto.MIMEHeader)
	partHeader.Add("Content-Type", "image/jpeg")

	for {
		if !webcam.IsOpen() {
			http.Error(w, "Missing camera", http.StatusServiceUnavailable)
			return
		}

		img, err := webcam.CaptureImage()
		if err != nil {
			http.Error(w, "Error capturing image: %v", http.StatusServiceUnavailable)
		}
		defer img.Close()

		resizedFrame := gocv.NewMat()

		gocv.Resize(*img, &resizedFrame, image.Point{X: 256, Y: 144}, 0, 0, gocv.InterpolationDefault)

		buf, err := gocv.IMEncode(gocv.JPEGFileExt, resizedFrame)
		if err != nil {
			log.Errorf("Error encoding frame: %v", err)
			continue
		}

		partWriter, _ := mimeWriter.CreatePart(partHeader)
		if _, err := partWriter.Write(buf.GetBytes()); err != nil {
			log.Error("Error while processing buffer")
		}

		// we want to record around 10 fps
		// this mean every second send 10 images
		// Let's assume reading, encoding and  writing do not consume any resources
		// If we sleep for 1/10 of a second we roughly approximate 10 fps
		time.Sleep(200 * time.Millisecond)
	}
}

// Create Camera using POST request
// TODO: Add device number option
func CreateCamera(w http.ResponseWriter, r *http.Request) {
	// Trigger singleton to instanciate camera
	utils.CreateOrGetCamera()
}

// Delete Camera using DELETE request
// TODO: Add device number option
func DeleteCamera(w http.ResponseWriter, r *http.Request) {
	utils.DeleteCamera()
}
