package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/utils"
	"github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/gin-gonic/gin"
)

type InteractHandler struct {
	port string
}

func NewInteractHandler(port string) *InteractHandler {
	return &InteractHandler{
		port: port,
	}
}

// Check Like godoc
//
//	@Summary		Check Like
//	@Description	Check Like with request and token
//	@Tags			interact
//	@Accept			json
//	@Produce		json
//	@Secure			BearerAuth
//	@Param			request	body		models.LikeRequest	true	"Like Check Request"
//	@Success		200		{object}	models.SwaggerLikeResp
//	@Failure		404		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/interact/check-like [post]
func (iH *InteractHandler) CheckLikeHandler(c *gin.Context) {
	server, client, input, err := utils.LikeRepoAndGetRequest(c, iH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	resp, err := client.CheckLike(ctx, input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	if resp.Valid == false {
		models.Response(c, http.StatusNotFound, false)
		return
	}

	models.Response(c, http.StatusOK, true)
}

// Add Like godoc
//
//	@Summary		Add Like
//	@Description	Add Like with request and token
//	@Tags			interact
//	@Accept			json
//	@Produce		json
//	@Secure			BearerAuth
//	@Param			request	body		models.LikeRequest	true	"Like Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/interact/add-like [post]
func (iH *InteractHandler) AddLikeHandler(c *gin.Context) {
	server, client, input, err := utils.LikeRepoAndGetRequest(c, iH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	_, err = client.AddLike(ctx, input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusCreated, "Add Success")
}

// Delete Like godoc
//
//	@Summary		Delete Like
//	@Description	Delete Like with request and token
//	@Tags			interact
//	@Accept			json
//	@Produce		json
//	@Secure			BearerAuth
//	@Param			request	body		models.LikeRequest	true	"Like Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/interact/delete-like [delete]
func (iH *InteractHandler) RemoveLikeHandler(c *gin.Context) {
	server, client, input, err := utils.LikeRepoAndGetRequest(c, iH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	_, err = client.RemoveLike(ctx, input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Delete Success")
}

// Add Comment godoc
//
//	@Summary		Add Comment
//	@Description	Add Comment with request and token
//	@Tags			interact
//	@Accept			json
//	@Produce		json
//	@Secure			BearerAuth
//	@Param			request	body		models.CommentRequest	true	"Comment Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/interact/add-comment [post]
func (iH *InteractHandler) AddCommentHandler(c *gin.Context) {
	server, client, err := repository.PostRepo(iH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	var req models.CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	input := &api.AddCommentReq{
		UserId:  req.UserId,
		PostId:  req.PostId,
		Comment: req.Comment,
	}

	_, err = client.AddComment(ctx, input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusCreated, "Comment Created")
}

// Delete Comment godoc
//
//	@Summary		Delete Comment
//	@Description	Delete Comment with request and token
//	@Tags			interact
//	@Accept			json
//	@Produce		json
//	@Secure			BearerAuth
//	@Param			id	query		string	true	"Id Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/interact/delete-comment [delete]
func (iH *InteractHandler) DeleteCommentHandler(c *gin.Context) {
	server, client, err := repository.PostRepo(iH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	input := &api.DeleteCommentReq{
		Id: c.Query("id"),
	}

	_, err = client.DeleteComment(ctx, input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Comment Deleted")
}
