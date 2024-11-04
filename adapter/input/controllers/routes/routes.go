package routes

import (
	controller "github.com/andamorim2206/hexagonal_1/adapter/input/controllers"
	"github.com/andamorim2206/hexagonal_1/adapter/output/news_http"
	"github.com/andamorim2206/hexagonal_1/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}