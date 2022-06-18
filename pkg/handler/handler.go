package handler

import (
	"github.com/Olexander753/constanta_test/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InutRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/id", h.getListByID)       //получение списка платежей по ID!
			list.GET("/email", h.getListByEmail) //проверка списка платажей по email!
		}

		item := api.Group("/item")
		{
			item.POST("/", h.postItem)        //создание платежа!
			item.GET("/check", h.getItemByID) //проверка статуса пдатежа по его ID!
			item.PUT("/", h.updateItemByID)   //изменение статуса платежа
		}
	}
	return router
}
