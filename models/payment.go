package models

import (
	"time"

	"gorm.io/gorm"
)

type PagoPendiente struct {
	gorm.Model

	IDReserva        uint
	Monto            float64
	Descripcion      string `gorm:"type:varchar(255)"`
	FechaVencimiento time.Time
	EstadoPago       string `gorm:"type:varchar(100);not null"`
}
