package handler

import (
	"go-cart-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProducts(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func (h *productHandler) GetProducts(ctx *gin.Context) {
	products, err := h.repo.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
