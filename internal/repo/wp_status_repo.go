package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

type WPStatusRepo interface {
	ListStatus() ([]model.WorkProgressStatus, error)
}

func NewWpStatusRepo() WPStatusRepo {
	return &impl.WPStatusRepoImpl{}
}
