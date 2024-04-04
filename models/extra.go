package models

import (
	"time"


)

type Alerta struct {
    ID                 uint
    IDUsuario          uint
    DescripcionAlerta  string    `gorm:"type:text"`
    FechaHoraCreacion  time.Time
    Usuario            Usuario   `gorm:"foreignKey:IDUsuario"`
}

type BitacoraAcceso struct {
    ID                 uint
    IDUsuario          uint
    FechaHoraAcceso    time.Time
    DireccionIP        string `gorm:"type:varchar(100);not null"`
    NavegadorUtilizado string `gorm:"type:varchar(255);not null"`
    Usuario            Usuario `gorm:"foreignKey:IDUsuario"`
}

type BitacoraActividad struct {
    ID                  uint
    IDUsuario           uint
    DescripcionActividad string    `gorm:"type:text"`
    FechaHoraActividad  time.Time
    Usuario             Usuario   `gorm:"foreignKey:IDUsuario"`
}

type Configuracion struct {
    ID                 uint
    NombreConfiguracion string `gorm:"type:varchar(100);not null;unique"`
    ValorConfiguracion  string `gorm:"type:text"`
}

type ComisionAgente struct {
    ID             uint
    IDAgente       uint
    IDTransaccion  uint
    MontoComision  float64
    FechaPago      time.Time
    Agente         AgenteInmobiliario `gorm:"foreignKey:IDAgente"`
    Transaccion    Transaccion        `gorm:"foreignKey:IDTransaccion"`
}

type SeguimientoCasa struct {
    ID                uint
    IDUsuario         uint
    IDCasa            uint
    FechaHoraVisita   time.Time
    InteresExpresado  string    `gorm:"type:text"`
    Usuario           Usuario   `gorm:"foreignKey:IDUsuario"`
    Casa              Casa      `gorm:"foreignKey:IDCasa"`
}

type HistorialPrecio struct {
    ID          uint
    IDCasa      uint
    Precio      float64
    FechaCambio time.Time
    Casa        Casa `gorm:"foreignKey:IDCasa"`
}

type SolicitudAdministrativa struct {
    ID                 uint
    IDUsuario          uint
    DescripcionSolicitud string    `gorm:"type:text"`
    EstadoSolicitud      string    `gorm:"type:varchar(100);not null"`
    FechaHoraSolicitud   time.Time
    Usuario              Usuario   `gorm:"foreignKey:IDUsuario"`
}

type Auditoria struct {
    ID                 uint
    Usuario             string    `gorm:"type:varchar(100);not null"`
    AccionRealizada     string    `gorm:"type:varchar(100);not null"`
    FechaHoraAccion     time.Time
    TablaAfectada       string    `gorm:"type:varchar(100);not null"`
    DetallesAdicionales string    `gorm:"type:text"`
}