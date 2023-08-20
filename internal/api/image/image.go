package image

import (
	"encoding/json"
	"github.com/thimovez/service/internal/api/middlewares"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type imageRouter struct {
	iImageRepo usecase.ImageRepo
}

func NewImageRoutes(handler *http.ServeMux, i usecase.ImageRepo, m *middlewares.Middleware) {
	r := &imageRouter{i}

	handler.HandleFunc("/upload-picture", m.AuthMiddleware(r.uploadPicture))
	handler.HandleFunc("/images", m.AuthMiddleware(r.getImages))
}

func (i *imageRouter) uploadPicture(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	userID := req.PostForm.Get("userID")
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
	uploadDir := "./static/images"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	url := req.URL
	imageURL := filepath.Join(url.String(), handler.Filename)
	imagePath := filepath.Join(uploadDir, handler.Filename)
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Unable to create the file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	image := entity.Image{
		ID:        "",
		UserID:    userID,
		ImagePath: imagePath,
		ImageURL:  imageURL,
	}

	err = i.iImageRepo.SaveImage(image)
	if err != nil {
		http.Error(w, "Error save file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func (i *imageRouter) getImages(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.Write([]byte("invalid method"))
		return
	}

	images, err := i.iImageRepo.GetImages()
	if err != nil {
		log.Fatalf("login service error %s", err)
	}

	marshal, err := json.Marshal(images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
