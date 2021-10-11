package handler

import (
	"go-cart-api/model"
	"go-cart-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	CreateTransaction(*gin.Context)
	GetTransactionsByUserId(*gin.Context)
	GetTransactionsById(*gin.Context)
	PayTransaction(*gin.Context)
	DeliverTransaction(*gin.Context)
}

type transactionHandler struct {
	repo repository.TransactionRepository
}

func NewTransactionHandler() TransactionHandler {

	return &transactionHandler{
		repo: repository.NewTransactionRepository(),
	}
}

func (h *transactionHandler) CreateTransaction(ctx *gin.Context) {
	var (
		transactionRequest model.TransactionCreateRequest
		totalPrice         uint64
	)

	if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItemRepo := repository.NewCartItemRepository()
	cartItems, err := cartItemRepo.GetItemsByIds(transactionRequest.CartItemIds)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not signed in"})
		return
	}

	for _, cartItem := range cartItems {
		totalPrice += cartItem.Product.Price * uint64(cartItem.Quantity)
	}

	transaction := model.Transaction{
		CartItems:  cartItems,
		State:      "pending",
		UserId:     uint(userId.(float64)),
		TotalPrice: uint(totalPrice),
	}

	transaction, err = h.repo.CreateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *transactionHandler) GetTransactionsByUserId(ctx *gin.Context) {
	userId, ok := ctx.Get("UserID")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not signed in"})
		return
	}

	transactions, err := h.repo.GetTransactionsByUserId(uint(userId.(float64)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, transactions)
}

func (h *transactionHandler) GetTransactionsById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.repo.GetTransactionsById(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *transactionHandler) PayTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.repo.GetTransactionsById(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if transaction.PaymentMethod == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "payment_method can not be empty"})
		return
	}

	transaction.State = "paid"
	transaction, err = h.repo.UpdateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *transactionHandler) DeliverTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.repo.GetTransactionsById(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction.State = "delivered"
	transaction.RemitAmount = transaction.TotalPrice
	transaction, err = h.repo.UpdateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add remit function to send the remit_amount
	// money to seller (assumming seller is is BlanjaAja)

	ctx.JSON(http.StatusOK, transaction)
}
