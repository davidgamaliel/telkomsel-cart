package handler

import (
	"go-cart-api/model"
	"go-cart-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartItemHandler interface {
	CreateItem(*gin.Context)
	GetItemsByUserId(*gin.Context)
}

type cartItemHandler struct {
	repo repository.CartItemRepository
}

func NewCartItemHandler() CartItemHandler {

	return &cartItemHandler{
		repo: repository.NewCartItemRepository(),
	}
}

func (h *cartItemHandler) CreateItem(ctx *gin.Context) {
	var cartItem model.CartItem
	userId, ok := ctx.Get("UserID")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not signed in"})
		return
	}

	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItem.UserId = uint(userId.(float64))

	cartItem, err := h.repo.CreateItem(cartItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, cartItem)
}

func (h *cartItemHandler) GetItemsByUserId(ctx *gin.Context) {
	userId, ok := ctx.Get("UserID")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not signed in"})
		return
	}

	cartItems, err := h.repo.GetItemsByUserId(uint(userId.(float64)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, cartItems)
}
