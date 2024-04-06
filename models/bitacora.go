package models

import (
	"time"
)


type BitacoraAcceso struct {
	ID_bitacora_acceso uint `gorm:"primaryKey;column:ID_bitacora_acceso"`
	IDUsuario          uint
	FechaHoraAcceso    time.Time
	DireccionIP        string  `gorm:"type:varchar(100);not null"`
	NavegadorUtilizado string  `gorm:"type:varchar(255);not null"`
	Usuario 			Usuario `gorm:"foreignKey:IDUsuario"`
	
}

type BitacoraActividad struct {
	ID_bitacora_actividad uint `gorm:"primaryKey;column:ID_bitacora_actividad"`
	IDUsuario             uint
	DescripcionActividad  string `gorm:"type:text"`
	FechaHoraActividad    time.Time
	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
}


type Auditoria struct {
	ID_auditoria        uint   `gorm:"primaryKey;column:ID_auditoria"`
	IDUsuario              string `gorm:"type:varchar(100);not null"`
	AccionRealizada     string `gorm:"type:varchar(100);not null"`
	FechaHoraAccion     time.Time
	TablaAfectada       string `gorm:"type:varchar(100);not null"`
	DetallesAdicionales string `gorm:"type:text"`
	Usuario 			Usuario `gorm:"foreignKey:IDUsuario"`
}
