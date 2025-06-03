package controller

import (
	"first-project/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LikeArticle(context *gin.Context) {
	articleID := context.Param("id")

	likeKey := "article:" + articleID + ":likes"
	global.RedisDB.Incr(likeKey).Err()
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func getArticleLikes(context *gin.Context) {
	articleID := context.Param("id")
	likeKey := "article:" + articleID + ":likes"
	global.RedisDB.Get(likeKey).Result()
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}
