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

type BitacoraAcceso struct {
	ID_bitacora_acceso uint `gorm:"primaryKey;column:ID_bitacora_acceso"`
	IDUsuario          uint
	FechaHoraAcceso    time.Time
	DireccionIP        string  `gorm:"type:varchar(100);not null"`
	NavegadorUtilizado string  `gorm:"type:varchar(255);not null"`
	Usuario            Usuario `gorm:"foreignKey:IDUsuario"`
}

type BitacoraActividad struct {
	ID_bitacora_actividad uint `gorm:"primaryKey;column:ID_bitacora_actividad"`
	IDUsuario             uint
	DescripcionActividad  string `gorm:"type:text"`
	FechaHoraActividad    time.Time
	Usuario               Usuario `gorm:"foreignKey:IDUsuario"`
}

type Configuracion struct {
	ID_configuracion    uint   `gorm:"primaryKey;column:ID_configuracion"`
	NombreConfiguracion string `gorm:"type:varchar(100);not null;unique"`
	ValorConfiguracion  string `gorm:"type:text"`
}

type ComisionAgente struct {
	ID_comision_agente uint `gorm:"primaryKey;column:ID_comision_agente"`
	IDAgente           uint
	IDTransaccion      uint
	MontoComision      float64
	FechaPago          time.Time
	Agente             AgenteInmobiliario `gorm:"foreignKey:IDAgente"`
	Transaccion        Transaccion        `gorm:"foreignKey:IDTransaccion"`
}

type SeguimientoCasa struct {
	ID_seguimiento_casa uint `gorm:"primaryKey;column:ID_seguimiento_casa"`
	IDUsuario           uint
	IDCasa              uint
	FechaHoraVisita     time.Time
	InteresExpresado    string  `gorm:"type:text"`
	Usuario             Usuario `gorm:"foreignKey:IDUsuario"`
	Casa                Casa    `gorm:"foreignKey:IDCasa"`
}

type HistorialPrecio struct {
	ID_historial_precio uint `gorm:"primaryKey;column:ID_historial_precio"`
	IDCasa              uint
	Precio              float64
	FechaCambio         time.Time
	Casa                Casa `gorm:"foreignKey:IDCasa"`
}

type SolicitudAdministrativa struct {
	ID_solicitud_administrativa uint `gorm:"primaryKey;column:ID_solicitud_administrativa"`
	IDUsuario                   uint
	DescripcionSolicitud        string `gorm:"type:text"`
	EstadoSolicitud             string `gorm:"type:varchar(100);not null"`
	FechaHoraSolicitud          time.Time
	Usuario                     Usuario `gorm:"foreignKey:IDUsuario"`
}

type Auditoria struct {
	ID_auditoria        uint   `gorm:"primaryKey;column:ID_auditoria"`
	Usuario             string `gorm:"type:varchar(100);not null"`
	AccionRealizada     string `gorm:"type:varchar(100);not null"`
	FechaHoraAccion     time.Time
	TablaAfectada       string `gorm:"type:varchar(100);not null"`
	DetallesAdicionales string `gorm:"type:text"`
}
