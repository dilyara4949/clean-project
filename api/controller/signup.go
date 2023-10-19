package controller

import (
	"net/http"

	"github.com/dilyara4949/clean-project/bootstrap"
	"github.com/dilyara4949/clean-project/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	us, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if us.Email != "" {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

// package controller

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/dilyara4949/clean-project/bootstrap"
// 	"github.com/dilyara4949/clean-project/domain"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"golang.org/x/crypto/bcrypt"
// )

// type SignupController struct {
// 	SignupUsecase domain.SignupUsecase
// 	Env *bootstrap.Env
// }

// func (sc *SignupController) Signup(c *gin.Context) {
// 	var request domain.SignupRequest

// 	err := c.ShouldBind(&request)
// 	if err != nil {
		
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User already exists with the given email1"})
// 	}
	
// 	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
// 	if err != nil {
// 		fmt.Println(request.Email)
// 		fmt.Println(request.Name)
// 		fmt.Println(request.Password)
// 		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email2"})
// 	}

// 	encryptedPassword, err := bcrypt.GenerateFromPassword(
// 		[]byte(request.Password),
// 		bcrypt.DefaultCost,
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User already exists with the given email3"})
// 		return
// 	}

// 	request.Password = string(encryptedPassword)

// 	user := domain.User{
// 		ID:       primitive.NewObjectID(),
// 		Name:     request.Name,
// 		Email:    request.Email,
// 		Password: request.Password,
// 	}
// 	fmt.Println(user)
// 	err = sc.SignupUsecase.Create(c, &user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User already exists with the given emai4l"})
// 		return
// 	}

// 	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User already exists with the given em5ail"})
// 		return
// 	}

// 	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User already exists with 6the given email"})
// 		return
// 	}

// 	signupResponse := domain.SignupResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}

// 	c.JSON(http.StatusOK, signupResponse)
// }