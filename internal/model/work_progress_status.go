package model

type WorkProgressStatus struct {
	Base
	StatusName string `json:"status_name" gorm:"column:status_name;type:varchar(255);not null;"`
}
