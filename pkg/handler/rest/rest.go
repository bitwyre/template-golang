package rest

import "github.com/bitwyre/template-golang/pkg/service"

type Rest struct {
	*service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		Service: service,
	}
}
