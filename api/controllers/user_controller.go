package controllers

import (
	"net/http"
	"strconv"
	"eform-gateway/api/services"
	"eform-gateway/constants"
	"eform-gateway/lib"
	"eform-gateway/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := u.service.GetOneUser(uint(id))

	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		u.logger.Zap.Error(err)
	}
	c.JSON(200, gin.H{"data": users})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.Bind(&user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.WithTrx(trxHandle).CreateUser(user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{"data": "user created"})
}

// SaveUserWOTrx saves the user without transaction for comparision
func (u UserController) SaveUserWOTrx(c *gin.Context) {
	user := models.User{}

	
	if err := c.Bind(&user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error bind JSON": err.Error(),
		})
		return
	}

	if err := u.service.CreateUser(user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	paramID := c.Param("id")

	if err := c.Bind(&user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.UpdateUser(uint(id), user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user updated"})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.DeleteUser(uint(id)); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
