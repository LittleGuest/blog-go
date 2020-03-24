package model

import (
	"time"
)

type FlywaySchemaHistory struct {
	InstalledRank int       `json:"installed_rank" gorm:"type:int;not null;primary_key"`
	Version       string    `json:"version" gorm:"type:varchar(50);null"`
	Description   string    `json:"description" gorm:"type:varchar(200);not null"`
	Type          string    `json:"type" gorm:"type:varchar(20);not null"`
	Script        string    `json:"script" gorm:"type:varchar(1000);not null"`
	Checksum      int       `json:"checksum" gorm:"type:int;null"`
	InstalledBy   string    `json:"installed_by" gorm:"type:varchar(100);not null"`
	InstalledOn   time.Time `json:"installed_on" gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null"`
	ExecutionTime int       `json:"execution_time" gorm:"type:int;not null"`
	Success       int       `json:"success" gorm:"type:int;not null"`
}
