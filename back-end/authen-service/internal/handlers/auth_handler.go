package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"github.com/billzayy/social-media/back-end/authen-service/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(us *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: us,
	}
}

func (aH *AuthHandler) RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err)
		return
	}

	resp, err := aH.AuthService.RegisterService(req)

	if err != nil || !resp {
		models.ResponseUser(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.ResponseUser(c, http.StatusCreated, "Account Created !")
}

func (aH *AuthHandler) LoginHandler(c *gin.Context) {
	userName := c.Query("userName")
	password := c.Query("password")

	res, cookie, err := aH.AuthService.LoginService(userName, password)

	if err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(c.Writer, &cookie) // Set Cookie

	// Return Status OK
	models.ResponseUser(c, http.StatusOK, res)
}

// This function will return a new AccessToken and RefreshToken
// with requires the client send the old Refresh Token to run this API
func (aH *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	// Get a Refresh Token
	getToken, err := c.Cookie("jwt")

	if err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err.Error())
		return
	}

	userData, newToken, expires, cookie, err := aH.AuthService.RefreshTokenService(getToken)

	if err != nil {
		models.ResponseUser(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &cookie)

	// Return like LoginHandler
	models.ResponseUser(c, http.StatusOK, models.RefreshTokenResp{
		User:      userData,
		Token:     newToken,
		Type:      "Bearer",
		ExpiresIn: expires,
	})
}
