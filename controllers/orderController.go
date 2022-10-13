package controllers

import (
	"net/http"
	"restapi/database"
	"restapi/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Get Orders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	var (
		orders []models.Orders
	)

	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"orders": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
		"count":  len(orders),
	})
}

// Get Detauk Orders godoc
// @Summary Get Detail orders
// @Description Get Detail orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /orders/:id [get]
func GetOrderDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusAccepted, gin.H{
		"id": id,
	})
	db := database.GetDB()
	var detailOrder models.Orders

	err := db.Preload("Items").First(&detailOrder, id).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"result": detailOrder,
	})
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var newOrder models.Orders
	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	newOrder.Ordered_at = time.Now()
	err = db.Create(&newOrder).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": newOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {

	db := database.GetDB()
	id := ctx.Param("id")
	// var requestBody models.Orders
	var detailValue models.Orders
	var updateValue models.Orders

	err := ctx.ShouldBindJSON(&updateValue)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"result": nil,
			"err":    "Wrong body format",
		})
		return
	}

	errDetail := db.First(&detailValue, id).Error
	if errDetail != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"result": nil,
			"err":    "Data not found",
		})
		return
	}

	errUpdate := db.Model(&detailValue).Updates(models.Orders{
		Customer_name: updateValue.Customer_name,
		Items:         updateValue.Items,
	}).Error

	errUpdateAssoc := db.Model(&detailValue).Association("Items").Replace(&updateValue.Items)

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    errUpdate,
		})
		return
	}

	if errUpdateAssoc != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    errUpdateAssoc,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": updateValue,
	})

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")

	errDelete := db.Where("order_id= ?", id).Delete(&models.Orders{}).Error
	if errDelete != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"result": nil,
			"err":    errDelete,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "deleted",
	})

}
