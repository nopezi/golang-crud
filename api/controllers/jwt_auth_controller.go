package controllers

import (
	"net/http"
	"clean-gin-template/api/services"
	"clean-gin-template/lib"
	"clean-gin-template/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     services.JWTAuthService
	userService services.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service services.JWTAuthService,
	userService services.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	user := &models.User{}

	if err := c.Bind(&user); err != nil {
		jwt.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error bind JSON": err.Error(),
		})
	}
	
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	result, err := jwt.userService.GetOneUserEmail(user.Email)
	

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	token := jwt.service.CreateToken(result)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		jwt.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error bind JSON": err.Error(),
		})
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)

	if err := jwt.userService.CreateUser(user); err != nil {
		jwt.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Register Successfully",
	})
}
