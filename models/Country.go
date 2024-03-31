package models

import "gorm.io/gorm"

type Pais struct {
	gorm.Model

	NombrePais string `gorm:"type:varchar(100);not null"`
	// Ciudades   []Ciudad // Relationship with Ciudades
}
