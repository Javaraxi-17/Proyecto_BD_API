package models

import (
	"time"


)

type Reserva struct {
    ID              uint
    IDUsuario       uint
    IDCasa          uint
    FechaHoraInicio time.Time
    FechaHoraFin    time.Time
    EstadoReserva   string    `gorm:"type:varchar(100);not null"`
    PrecioTotal     float64
    Usuario         Usuario   `gorm:"foreignKey:IDUsuario"`
    Casa            Casa      `gorm:"foreignKey:IDCasa"`
    Comentarios     []Comentario `gorm:"foreignKey:IDReserva"`
}

type ReservaHistorica struct {
    ID              uint
    IDUsuario       uint
    IDCasa          uint
    FechaHoraInicio time.Time
    FechaHoraFin    time.Time
    EstadoReserva   string    `gorm:"type:varchar(100);not null"`
    PrecioTotal     float64
    Usuario         Usuario   `gorm:"foreignKey:IDUsuario"`
    Casa            Casa      `gorm:"foreignKey:IDCasa"`
}
