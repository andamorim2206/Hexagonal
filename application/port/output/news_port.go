package output

import (
	"github.com/andamorim2206/hexagonal_1/application/domain"
	"github.com/andamorim2206/hexagonal_1/configuration/rest_err"
)

type GetNewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}