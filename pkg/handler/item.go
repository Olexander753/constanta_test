package handler

import (
	"log"
	"net/http"

	test "github.com/Olexander753/constanta_test"
	"github.com/gin-gonic/gin"
)

type getResponse struct {
	Data_ test.Transactions `json:"data"`
}

func (h *Handler) postItem(c *gin.Context) {
	var input test.Transaction

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 1: ", err)
	}

	id, err := h.services.TransactionItem.CreateTransaction(input)
	if err != nil {
		log.Fatal("error insert: ", err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "траназакция успешно создана",
	})
}

func (h *Handler) getItemByID(c *gin.Context) {
	var input test.Transactions
	var list test.Transactions

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 3: ", err)
	}

	list, err := h.services.TransactionItem.GetTransaction(input.ID)

	if err != nil {
		log.Fatal("error get list 1: ", err)
	}

	c.JSON(http.StatusOK, getResponse{
		Data_: list,
	})
}

func (h *Handler) updateItemByID(c *gin.Context) {
	var input test.Action

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 6: ", err)
	}

	err := h.services.TransactionItem.UpdateTransaction(input.ID, input.Discription)

	if err != nil {
		c.JSON(http.StatusOK, "Ошибка")
	} else {
		c.JSON(http.StatusOK, "OK")
	}

}
