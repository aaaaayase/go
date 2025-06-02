package controller

import (
	"first-project/global"
	"first-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateExchangeRate(context *gin.Context) {
	var exchangeRate model.ExchangeRate
	err := context.ShouldBind(&exchangeRate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	err = global.Db.AutoMigrate(&exchangeRate)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = global.Db.Create(&exchangeRate).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, exchangeRate)
}
