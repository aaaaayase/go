package controller

import (
	"first-project/global"
	"first-project/model"
	"first-project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user.Password = hashedPwd
	token, err := utils.GenerateJWT(user.UserName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = global.Db.AutoMigrate(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if global.Db.Create(&user).Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Login(context *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := context.ShouldBind(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User

	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !utils.CheckPasswword(user.Password, input.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
	}

	token, err := utils.GenerateJWT(user.UserName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
