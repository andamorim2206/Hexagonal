package controller

import (
	"net/http"

	"github.com/andamorim2206/hexagonal_1/adapter/input/controllers/model/request"
	"github.com/andamorim2206/hexagonal_1/application/domain"
	"github.com/andamorim2206/hexagonal_1/application/port/input"
	"github.com/andamorim2206/hexagonal_1/configuration/logger"
	"github.com/andamorim2206/hexagonal_1/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *newsController {
	return &newsController{newsUseCase}
}
func (nc *newsController) GetNews(c *gin.Context) {

	logger.Info("Init GetNews controller api")

	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	
	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From: newsRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.newsUseCase.GetNewsServices(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
