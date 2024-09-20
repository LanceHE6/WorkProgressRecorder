package service

import (
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

type WPStatusService interface {
	GetStatus() *pkg.Response
}

func NewWpStatusService() WPStatusService {
	return &impl.WPStatusServiceImpl{}
}
