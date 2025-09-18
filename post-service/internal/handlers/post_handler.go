package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/post-service/internal/models"
	"github.com/billzayy/social-media/post-service/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostService *services.PostService
}

func NewPostHandler(ur *services.PostService) *PostHandler {
	return &PostHandler{
		PostService: ur,
	}
}

func (pH *PostHandler) CreatePostHandler(c *gin.Context) {
	var req models.AddPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	if req.Content == "" {
		models.Response(c, http.StatusBadRequest, "Content can not be empty!")
		return
	}

	valid, err := pH.PostService.CreatePost(req)

	if err != nil && err.Error() != "content can not empty" {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	if valid == false {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	models.Response(c, http.StatusCreated, "Post Created!")
}

func (pH *PostHandler) GetPostHandler(c *gin.Context) {
	resp, err := pH.PostService.GetPost()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, resp)
}

func (pH *PostHandler) DeletePostHandler(c *gin.Context) {
	postId := c.Query("id")

	err := pH.PostService.DeletePost(postId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Delete Successful!")
}
