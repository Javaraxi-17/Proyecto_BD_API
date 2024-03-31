package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model

	Nombre string `gorm:"type:varchar(100);not null"`
}
