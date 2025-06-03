package controller

import (
	"first-project/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func LikeArticle(context *gin.Context) {
	articleID := context.Param("id")

	likeKey := "article:" + articleID + ":likes"
	global.RedisDB.Incr(likeKey).Err()
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetArticleLikes(context *gin.Context) {
	articleID := context.Param("id")
	likeKey := "article:" + articleID + ":likes"
	likes, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"likes": likes})
}
