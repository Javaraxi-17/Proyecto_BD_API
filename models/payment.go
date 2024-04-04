package models

import (
	"time"


)


type PagoPendiente struct {
    ID              uint
    IDReserva       uint
    Monto           float64
    Descripcion     string    `gorm:"type:varchar(255)"`
    FechaVencimiento time.Time
    EstadoPago      string    `gorm:"type:varchar(100);not null"`
    Reserva         Reserva   `gorm:"foreignKey:IDReserva"`
}