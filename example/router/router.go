package router

import (
	"golte/controller"
	"golte/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	VideoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

type Product struct {
	Id          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
}

func Router(server *gin.Engine) {
	indexRoutes := server.Group("/")
	{
		indexRoutes.GET("", VideoController.ShowAll)
	}

	server.POST("/products/product", func(ctx *gin.Context) {
		var product Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Product Input Not Valid!!"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Product Input Valid!!"})
	})
}
