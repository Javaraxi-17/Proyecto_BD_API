package models

import (
	"time"

	"gorm.io/gorm"
)

type Alerta struct {
	gorm.Model

	IDUsuario         uint
	DescripcionAlerta string `gorm:"type:text"`
	FechaHoraCreacion time.Time
}

type BitacoraAcceso struct {
	gorm.Model

	IDUsuario          uint
	FechaHoraAcceso    time.Time
	DireccionIP        string `gorm:"type:varchar(100);not null"`
	NavegadorUtilizado string `gorm:"type:varchar(255);not null"`
}

type BitacoraActividad struct {
	gorm.Model

	IDUsuario            uint
	DescripcionActividad string `gorm:"type:text"`
	FechaHoraActividad   time.Time
}

type Configuracion struct {
	gorm.Model

	NombreConfiguracion string `gorm:"type:varchar(100);not null"`
	ValorConfiguracion  string `gorm:"type:text"`
}

type ComisionAgente struct {
	gorm.Model

	IDAgente      uint
	IDTransaccion uint
	MontoComision float64
	FechaPago     time.Time
}

type SeguimientoCasa struct {
	gorm.Model

	IDUsuario        uint
	IDCasa           uint
	FechaHoraVisita  time.Time
	InteresExpresado string `gorm:"type:text"`
}

type HistorialPrecio struct {
	gorm.Model

	IDCasa      uint
	Precio      float64
	FechaCambio time.Time
}

type SolicitudAdministrativa struct {
	gorm.Model

	IDUsuario            uint
	DescripcionSolicitud string `gorm:"type:text"`
	EstadoSolicitud      string `gorm:"type:varchar(100);not null"`
	FechaHoraSolicitud   time.Time
}

type Auditoria struct {
	gorm.Model

	Usuario             string `gorm:"type:varchar(100);not null"`
	AccionRealizada     string `gorm:"type:varchar(100);not null"`
	FechaHoraAccion     time.Time
	TablaAfectada       string `gorm:"type:varchar(100);not null"`
	DetallesAdicionales string `gorm:"type:text"`
}
