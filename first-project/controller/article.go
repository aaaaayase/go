package controller

import (
	"first-project/global"
	"first-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(context *gin.Context) {
	var article model.Article
	err := context.ShouldBind(&article)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = global.Db.AutoMigrate(&article)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = global.Db.Create(&article).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"article": article})
}

func GetArticle(context *gin.Context) {

	var articles []model.Article
	err := global.Db.Find(&articles).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"articles": articles})
}

func GetArticleByID(context *gin.Context) {
	id := context.Param("id")
	var article model.Article
	err := global.Db.Where("id = ?", id).First(&article).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"article": article})

}
