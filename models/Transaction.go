package models

import (
	"time"
)

type Transaccion struct {
	ID_transaccion    uint `gorm:"primaryKey;column:ID_transaccion"`
	IDUsuario         uint
	TipoTransaccion   string `gorm:"type:varchar(100);not null"`
	Monto             float64
	FechaHora         time.Time
	MetodoPago        string           `gorm:"type:varchar(100)"`
	EstadoTransaccion string           `gorm:"type:varchar(100);not null"`
	// Comisiones        []ComisionAgente `gorm:"foreignKey:IDTransaccion"`
}
