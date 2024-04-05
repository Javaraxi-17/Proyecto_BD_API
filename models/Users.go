package models

import (
	"time"
)

type Usuario struct {
	ID             uint   `gorm:"primaryKey"`
	Nombre         string `gorm:"type:varchar(100);not null"`
	Correo         string `gorm:"type:varchar(100);not null;unique"`
	Contrasena     string `gorm:"type:varchar(100);not null"`
	Fecha_registro time.Time
	Roles          []Rol     `gorm:"many2many:usuarios_roles;"`
	Reservas       []Reserva `gorm:"foreignKey:IDUsuario"`
	// Alertas         []Alerta            `gorm:"foreignKey:IDUsuario"`
	BitacorasAcceso    []BitacoraAcceso          `gorm:"foreignKey:IDUsuario"`
	BitacorasActividad []BitacoraActividad       `gorm:"foreignKey:IDUsuario"`
	SolicitudesAdmin   []SolicitudAdministrativa `gorm:"foreignKey:IDUsuario"`
}
