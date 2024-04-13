package models

import (
	"time"
)

type Alerta struct {
	ID_alerta         uint `gorm:"primaryKey;column:ID_alerta"`
	IDUsuario         uint
	DescripcionAlerta string `gorm:"type:text"`
	FechaHoraCreacion time.Time
}


type Configuracion struct {
	ID_configuracion    uint   `gorm:"primaryKey;column:ID_configuracion"`
	NombreConfiguracion string `gorm:"type:varchar(100);not null;unique"`
	ValorConfiguracion  string `gorm:"type:text"`
}



type HistorialPrecio struct {
	ID_historial_precio uint `gorm:"primaryKey;column:ID_historial_precio"`
	IDCasa              uint
	Precio              float64
	FechaCambio         time.Time
}

type SolicitudAdministrativa struct {
	ID_solicitud_administrativa uint `gorm:"primaryKey;column:ID_solicitud_administrativa"`
	IDUsuario                   uint
	DescripcionSolicitud        string `gorm:"type:text"`
	EstadoSolicitud             string `gorm:"type:varchar(100);not null"`
	FechaHoraSolicitud          time.Time
}



