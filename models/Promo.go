package models

import (
	"time"

	"gorm.io/gorm"
)

type Promocion struct {
	gorm.Model

	CodigoPromocion string `gorm:"type:varchar(100);not null"`
	Descuento       float64
	FechaInicio     time.Time
	FechaFin        time.Time
	Condiciones     string `gorm:"type:text"`
	// Reservas        []Reserva `gorm:"many2many:aplicaciones_promociones;"`
}

type AplicacionPromocion struct {
	gorm.Model

	IDReserva         uint
	IDPromocion       uint
	DescuentoAplicado float64
}
