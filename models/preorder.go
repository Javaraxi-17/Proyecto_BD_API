package models

import (
	"time"
)

type Reserva struct {
	ID_reserva            uint `gorm:"primaryKey;column:ID_reserva"`
	IDUsuario             uint
	IDCasa                uint
	FechaHoraInicio       time.Time
	FechaHoraFin          time.Time
	EstadoReserva         string `gorm:"type:varchar(100);not null"`
	PrecioTotal           float64
	// PagosPendientes       []PagoPendiente       `gorm:"foreignKey:IDReserva"`
	// AplicacionesPromocion []AplicacionPromocion `gorm:"foreignKey:IDReserva"`
}

type ReservaHistorica struct {
	ID_reserva_historica uint `gorm:"primaryKey;column:ID_reserva_historica"`
	IDUsuario            uint
	IDCasa               uint
	FechaHoraInicio      time.Time
	FechaHoraFin         time.Time
	EstadoReserva        string `gorm:"type:varchar(100);not null"`
	PrecioTotal          float64

}
