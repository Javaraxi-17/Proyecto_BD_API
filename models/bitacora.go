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
	
}

type BitacoraActividad struct {
	ID_bitacora_actividad uint `gorm:"primaryKey;column:ID_bitacora_actividad"`
	IDUsuario             uint
	DescripcionActividad  string `gorm:"type:text"`
	FechaHoraActividad    time.Time
}


type Auditoria struct {
	ID_auditoria        uint   `gorm:"primaryKey;column:ID_auditoria"`
	IDUsuario           uint
	AccionRealizada     string 
	FechaHoraAccion     time.Time
	TablaAfectada       string 
	DetallesAdicionales string `gorm:"type:text"`
}
