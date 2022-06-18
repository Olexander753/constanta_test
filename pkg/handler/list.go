package handler

import (
	"log"
	"net/http"

	test "github.com/Olexander753/constanta_test"
	"github.com/gin-gonic/gin"
)

type getListResponse struct {
	Data []test.Transactions `json:"data"`
}

func (h *Handler) getListByID(c *gin.Context) {
	var input test.Transactions
	var list []test.Transactions

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 3: ", err)
	}

	list, err := h.services.TransactionsList.GetListByID(input.User_id)

	if err != nil {
		log.Fatal("error get list 1: ", err)
	}

	c.JSON(http.StatusOK, getListResponse{
		Data: list,
	})
}

func (h *Handler) getListByEmail(c *gin.Context) {
	var input test.Transactions
	var list []test.Transactions

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 4: ", err)
	}

	list, err := h.services.TransactionsList.GetListByEmail(input.User_email)

	if err != nil {
		log.Fatal("error get list 2: ", err)
	}

	c.JSON(http.StatusOK, getListResponse{
		Data: list,
	})
}
