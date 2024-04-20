package comment

import (
	"encoding/json"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/comment"
	"net/http"
)

type commentRouter struct {
	iCommentService *comment.UseCaseComment
}

func NewCommentRoutes(handler *http.ServeMux, c *comment.UseCaseComment) {
	r := &commentRouter{
		iCommentService: c,
	}

	handler.HandleFunc("/comment/create", r.createComment)
}

func (u *commentRouter) createComment(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		w.Write([]byte("invalid method"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var c entity.Comment
	err := decoder.Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//c.UserID = req.PostForm.Get("userID")

	err = u.iCommentService.CreateComment(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
