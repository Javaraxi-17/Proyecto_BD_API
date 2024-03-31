package models

import "gorm.io/gorm"

type Ciudad struct {
	gorm.Model

	NombreCiudad string   `gorm:"type:varchar(100);not null"`
	PaisID       uint     // Foreign key to Pais
	Barrios      []Barrio `gorm:"many2many:ciudad_barrio;"`
}

type Barrio struct {
	gorm.Model

	NombreBarrio string   `gorm:"type:varchar(100);not null"`
	Ciudades     []Ciudad `gorm:"many2many:ciudad_barrio;"`
}
