package controllers

import (
	"Tugas2golang-gin/apimodels"
	"net/http"

	"github.com/gin-gonic/gin"
	
)
func CreateOrder(ctx *gin.Context) {
	var res apimodels.Response
	var req apimodels.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := services.SaveOrder(req)
	if err != nil {
		res.Status = "Create Ordel Gagal"
		res.ResponseCode = "400"

	}
	ctx.JSON(http.StatusOK, res)
}
func GetOrderAll(ctx *gin.Context) {

}

func UpdateOrder(ctx *gin.Context) {

}
func DeleteOrder(ctx *gin.Context) {

}
