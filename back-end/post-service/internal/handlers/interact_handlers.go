package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/billzayy/social-media/back-end/post-service/internal/services"
	"github.com/gin-gonic/gin"
)

type InteractHandler struct {
	InteractService *services.InteractService
}

func NewInteractHandler(iS *services.InteractService) *InteractHandler {
	return &InteractHandler{
		InteractService: iS,
	}
}

func (iH *InteractHandler) CheckLikeHandler(c *gin.Context) {
	var req models.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := iH.InteractService.CheckLikeOnPostService(req.UserId, req.PostId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data == false {
		models.Response(c, http.StatusNotFound, "not found")
		return
	}

	models.Response(c, http.StatusOK, true)
}

func (iH *InteractHandler) AddLikeHandler(c *gin.Context) {
	var req models.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadGateway, err.Error())
		return
	}

	err := iH.InteractService.AddLikeService(req.UserId, req.PostId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusCreated, "Liked!")
}

func (iH *InteractHandler) RemoveLikeHandler(c *gin.Context) {
	var req models.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	err := iH.InteractService.RemoveLikeService(req.UserId, req.PostId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Removed!")
}

func (iH *InteractHandler) AddCommentHandler(c *gin.Context) {
	var req models.CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	err := iH.InteractService.AddCommentService(req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}
	models.Response(c, http.StatusCreated, "Added Comment!")
}

func (iH *InteractHandler) DeleteCommentHandler(c *gin.Context) {
	var req models.DeleteCommentReq

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	err := iH.InteractService.DeleteCommentService(req)

	if err != nil && err.Error() == "not found" {
		models.Response(c, http.StatusNotFound, "Not found comment")
		return
	}

	models.Response(c, http.StatusOK, "Deleted")
}
