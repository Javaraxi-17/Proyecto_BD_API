package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaccion struct {
	gorm.Model

	IDUsuario         uint
	TipoTransaccion   string `gorm:"type:varchar(100);not null"`
	Monto             float64
	FechaHora         time.Time
	MetodoPago        string `gorm:"type:varchar(100)"`
	EstadoTransaccion string `gorm:"type:varchar(100);not null"`
	// Comisiones        []ComisionAgente // Relationship with Comisiones_Agentes
}
