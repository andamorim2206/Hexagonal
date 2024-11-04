package service

import (
	"fmt"

	"github.com/andamorim2206/hexagonal_1/application/domain"
	"github.com/andamorim2206/hexagonal_1/application/port/output"
	"github.com/andamorim2206/hexagonal_1/configuration/logger"
	"github.com/andamorim2206/hexagonal_1/configuration/rest_err"
)

type newsService struct{
	newsPort output.GetNewsPort
}

func NewNewsService(newsPort output.GetNewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsServices(
	new domain.NewsReqDomain,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Init getNewsService function, subject=%s, from=%s", new.Subject, new.From))
	
	newsDomainResponse, err :=  ns.newsPort.GetNewsPort(new)
	
	return newsDomainResponse, err 
}
