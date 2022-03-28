package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddOrder(ctx *gin.Context) {
	db := database.GetDB()

	//inputData := map[string]interface{}{}
	order := models.Order{}
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	//fmt.Println(order)
	err = db.Create(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"message": "sukses order"})

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	//inputData := map[string]interface{}{}
	orderId := ctx.Param("orderId")

	if orderId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": "mohon masukan order id"})
		return
	}

	if _, err := strconv.Atoi(orderId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": "mohon masukan order id dalam angka"})
		return
	}

	err := db.Exec("DELETE FROM items where order_id =?", orderId).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}
	err = db.Exec("DELETE FROM orders where order_id =?", orderId).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"message": fmt.Sprintf("sukses delete order with id %s", orderId)})

}

func EditOrder(ctx *gin.Context) {
	db := database.GetDB()

	//inputData := map[string]interface{}{}
	orderIdParam := ctx.Param("orderId")

	if orderIdParam == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": "mohon masukan order id"})
		return
	}

	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": "mohon masukan order id dalam angka"})
		return
	}

	order := models.Order{}
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	order.OrderID = uint(orderId)

	for _, item := range order.Items {

		err = db.Model(&item).Updates(models.Item{ItemCode: item.ItemCode, Description: item.Description, Quantity: item.Quantity}).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
			return
		}
	}
	//fmt.Println(order)
	err = db.Save(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"message": "sukses edit order"})

}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()

	//inputData := map[string]interface{}{}
	order := []models.Order{}

	//fmt.Println(order)
	err := db.Preload("Items").Find(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"data": &order})

}
