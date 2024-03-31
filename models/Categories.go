package models

import "gorm.io/gorm"

type CategoriaPropiedad struct {
	gorm.Model

	Nombre      string `gorm:"type:varchar(100);not null"`
	Descripcion string `gorm:"type:varchar(255)"`
	// Casas       []Casa // Relationship with Casas through CasaCategoria
}

type CasaCategoria struct {
	gorm.Model

	IDCasa      uint
	IDCategoria uint
}
