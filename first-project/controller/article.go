package controller

import (
	"encoding/json"
	"first-project/global"
	"first-project/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

var cacheKey = "articles"

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
	err = global.RedisDB.Del(cacheKey).Err()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"article": article})
}

func GetArticle(context *gin.Context) {

	cacheData, err := global.RedisDB.Get(cacheKey).Result()

	if err == redis.Nil {
		var articles []model.Article
		err := global.Db.Find(&articles).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		articleJson, err := json.Marshal(articles)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = global.RedisDB.Set(cacheKey, articleJson, 10*time.Minute).Err()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"articles": articles})

	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var articles []model.Article
		err := json.Unmarshal([]byte(cacheData), &articles)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, gin.H{"articles": articles})
	}

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
