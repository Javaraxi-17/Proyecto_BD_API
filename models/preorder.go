package models

import (
	"time"

	"gorm.io/gorm"
)

type Reserva struct {
	gorm.Model

	IDUsuario       uint
	IDCasa          uint
	FechaHoraInicio time.Time
	FechaHoraFin    time.Time
	EstadoReserva   string `gorm:"type:varchar(100);not null"`
	PrecioTotal     float64
	// Comentarios     []Comentario    `gorm:"foreignKey:IDReserva"` // Relationship with Comentarios
	// PagosPendientes []PagoPendiente // Relationship with Pagos_Pendientes
}

type ReservaHistorica struct {
	gorm.Model

	IDUsuario       uint
	IDCasa          uint
	FechaHoraInicio time.Time
	FechaHoraFin    time.Time
	EstadoReserva   string `gorm:"type:varchar(100);not null"`
	PrecioTotal     float64
}
