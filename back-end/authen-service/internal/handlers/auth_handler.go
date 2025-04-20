package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	AuthRepository *repositories.AuthRepository
}

func NewUserHandler(ur *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{
		AuthRepository: ur,
	}
}

func (aH *AuthHandler) RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err)
		return
	}

	err := aH.AuthRepository.Register(req)

	if err != nil {
		models.ResponseUser(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.ResponseUser(c, http.StatusCreated, "Account Created !")
}

func (aH *AuthHandler) LoginHandler(c *gin.Context) {
	userName := c.Query("userName")
	password := c.Query("password")

	hashed, checkData, err := aH.AuthRepository.CheckUserWithNameAndPass(userName, password)

	if err != nil || !checkData {
		models.ResponseUser(c, http.StatusBadRequest, err.Error())
		return
	}

	res, cookie, err := aH.AuthRepository.Login(userName, hashed)

	if err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err.Error())
		return
	}

	// Return Status OK
	models.ResponseUser(c, http.StatusOK, res)
	http.SetCookie(c.Writer, &cookie)
}

// This function will return a new AccessToken and RefreshToken
// with requires the client send the old Refresh Token to run this API
func (aH *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	// Get a Refresh Token
	getCookie, err := c.Cookie("jwt")

	if err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, newToken, cookie, err := aH.AuthRepository.RefreshToken(getCookie)

	if err != nil {
		models.ResponseUser(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Return like LoginHandler
	models.ResponseUser(c, http.StatusOK, models.UserToken{
		UserId: uuid.MustParse(userId),
		Token:  newToken,
		Type:   "Bearer",
	})

	http.SetCookie(c.Writer, &cookie)
}
