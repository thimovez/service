package image

import (
	"encoding/json"
	"fmt"
	"github.com/thimovez/service/internal/api/middlewares"
	"github.com/thimovez/service/internal/usecase"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type imageRouter struct {
	i usecase.ImageRepo
}

func NewImageRoutes(handler *http.ServeMux, i usecase.ImageRepo, m *middlewares.Middleware) {
	r := &imageRouter{i}

	handler.HandleFunc("/upload-picture", m.AuthMiddleware(r.uploadPicture))
	handler.HandleFunc("/images", m.AuthMiddleware(r.getImages))
}

func (i *imageRouter) uploadPicture(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	// Parse the incoming form file
	err := req.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := req.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file to a directory (you can change the path as needed)
	uploadDir := "../uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	filePath := filepath.Join(uploadDir, handler.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create the file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Unable to copy the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}

func (i *imageRouter) getImages(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodGet {
		w.Write([]byte("invalid method"))
		return
	}

	// TODO write service for GetImages
	images, err := i.i.GetImages()
	if err != nil {
		log.Fatalf("login service error %s", err)
	}

	marshal, err := json.Marshal(images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)

}
