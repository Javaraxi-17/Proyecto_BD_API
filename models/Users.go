package models

import (
	"time"
)

type Usuario struct {
	Id_usuario                 uint   `gorm:"primaryKey;column:ID_usuario"`
	Nombre                     string `gorm:"type:varchar(100);not null"`
	Correo                     string `gorm:"type:varchar(100);not null;unique"`
	Contrasena                 string `gorm:"type:varchar(100);not null"`
	Fecha_registro             time.Time
	Reservas                   []Reserva                 `gorm:"foreignKey:IDUsuario"`
	Roles                      []Rol                     `gorm:"many2many:usuarios_roles;foreignKey:Id_usuario;joinForeignKey:id_usuario;References:ID_rol;joinReferences:id_rol"`
	Comentarios                []Comentario              `gorm:"foreignKey:IDUsuario"`
	Transacciones              []Transaccion             `gorm:"foreignKey:IDUsuario"`
	ReservasHistoricas         []ReservaHistorica        `gorm:"foreignKey:IDUsuario"`
	Alertas                    []Alerta                  `gorm:"foreignKey:IDUsuario"`
	BitacorasAcceso            []BitacoraAcceso          `gorm:"foreignKey:IDUsuario"`
	BitacorasActividad         []BitacoraActividad       `gorm:"foreignKey:IDUsuario"`
	SeguimientosCasa           []SeguimientoCasa         `gorm:"foreignKey:IDUsuario"`
	SolicitudesAdministrativas []SolicitudAdministrativa `gorm:"foreignKey:IDUsuario"`
	Auditorias                 []Auditoria               `gorm:"foreignKey:IDUsuario"`
}
