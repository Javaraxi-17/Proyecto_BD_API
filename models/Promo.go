package models

import (
	"time"


)

type Promocion struct {
    ID              uint
    CodigoPromocion string    `gorm:"type:varchar(100);not null;unique"`
    Descuento       float64
    FechaInicio     time.Time
    FechaFin        time.Time
    Condiciones     string    `gorm:"type:text"`
    Aplicaciones    []AplicacionPromocion `gorm:"foreignKey:IDPromocion"`
}

type AplicacionPromocion struct {
    ID               uint
    IDReserva        uint
    IDPromocion      uint
    DescuentoAplicado float64
    Reserva          Reserva   `gorm:"foreignKey:IDReserva"`
    Promocion        Promocion `gorm:"foreignKey:IDPromocion"`
}