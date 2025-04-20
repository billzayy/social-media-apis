package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostRepository *repositories.PostRepository
}

func NewPostHandler(ur *repositories.PostRepository) *PostHandler {
	return &PostHandler{
		PostRepository: ur,
	}
}

func (pH *PostHandler) AddPostHandler(c *gin.Context) {
	var req models.AddPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.ResponsePost(c, http.StatusBadRequest, err)
		return
	}

	err := pH.PostRepository.AddPost(req)

	if err != nil {
		models.ResponsePost(c, http.StatusInternalServerError, err)
		return
	}

	models.ResponsePost(c, http.StatusCreated, "Post Created!")
}

func (pH *PostHandler) GetPostHandler(c *gin.Context) {
	resp, err := pH.PostRepository.GetPost()

	if err != nil {
		models.ResponsePost(c, http.StatusInternalServerError, err)
		return
	}

	models.ResponsePost(c, http.StatusOK, resp)
}

func (pH *PostHandler) DeletePostHandler(c *gin.Context) {
	postId := c.Query("id")

	err := pH.PostRepository.DeletePost(postId)

	if err != nil {
		models.ResponsePost(c, http.StatusInternalServerError, err)
		return
	}

	models.ResponsePost(c, http.StatusOK, "Delete Successful!")
}
