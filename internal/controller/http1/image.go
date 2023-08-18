package http1

import (
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"log"
	"net/http"
)

type imageRouter struct {
	i usecase.ImageRepo
}

func NewImageRoutes(handler *http.ServeMux, i usecase.ImageRepo) {
	r := &imageRouter{i}

	handler.HandleFunc("/upload-picture", r.uploadPicture)
	handler.HandleFunc("/images", r.getImages)
}

func (i *imageRouter) uploadPicture(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var image entity.Image
	err := decoder.Decode(&image)
	if err != nil {
		panic(err)
	}

	err = i.i.SaveImage(image)
	if err != nil {
		log.Fatalf("upload picture error %s", err)
	}
}

func (i *imageRouter) getImages(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.Write([]byte("invalid method"))
		return
	}

	err := i.i.GetImages()
	if err != nil {
		log.Fatalf("login service error %s", err)
	}
}
