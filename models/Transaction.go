package models

import (
	"time"


)

type Transaccion struct {
    ID              uint
    IDUsuario       uint
    TipoTransaccion string    `gorm:"type:varchar(100);not null"`
    Monto           float64
    FechaHora       time.Time
    MetodoPago      string    `gorm:"type:varchar(100)"`
    EstadoTransaccion string `gorm:"type:varchar(100);not null"`
    Usuario         Usuario   `gorm:"foreignKey:IDUsuario"`
    Comisiones      []ComisionAgente `gorm:"foreignKey:IDTransaccion"`
}
