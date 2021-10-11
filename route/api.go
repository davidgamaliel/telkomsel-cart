package route

import (
	"go-cart-api/handler"
	"go-cart-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	cartItemHandler := handler.NewCartItemHandler()
	transactionHandler := handler.NewTransactionHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Blanja Aja")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/users")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignIn)
	}

	productRoutes := apiRoutes.Group("/products")
	{
		productRoutes.GET("/", productHandler.GetProducts)
	}

	exclusiveRoutes := apiRoutes.Group("/_exclusive", middleware.AuthorizeJWT())

	cartItemExclusiveRoutes := exclusiveRoutes.Group("/cart-items")
	{
		cartItemExclusiveRoutes.POST("/", cartItemHandler.CreateItem)
		cartItemExclusiveRoutes.GET("/", cartItemHandler.GetItemsByUserId)
	}

	transactionExclusiveRoutes := exclusiveRoutes.Group("/transactions")
	{
		transactionExclusiveRoutes.POST("/", transactionHandler.CreateTransaction)
		transactionExclusiveRoutes.GET("/", transactionHandler.GetTransactionsByUserId)
		transactionExclusiveRoutes.GET("/:id", transactionHandler.GetTransactionsById)
		transactionExclusiveRoutes.PATCH("/:id/pay", transactionHandler.PayTransaction)
		transactionExclusiveRoutes.PATCH("/:id/deliver", transactionHandler.DeliverTransaction)
	}

	return r.Run(address)
}
