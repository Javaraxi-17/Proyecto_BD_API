package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model

	ID_usuario     int       `gorm:"primaryKey"`                 // Standard field for the primary key
	Nombre         string    `gorm:"type:varchar(100);not null"` // A regular string field
	Correo         string    `gorm:"type:varchar(100);not null"` // Changed "Correo_electronico" to "Correo"
	Contrasena     string    `gorm:"type:varchar(100);not null"`
	Fecha_registro time.Time // Changed to time.Time directly
	Roles          []Rol     `gorm:"many2many:usuarios_roles;"` // Relationship with Roles using many-to-many
	// Comentarios     []Comentario        `gorm:"foreignKey:IDUsuario"`      // Relationship with Comentarios
	// Reservas        []Reserva           // Relationship with Reservas
	// PagosPendientes []PagoPendiente     // Relationship with Pagos_Pendientes
	// Alertas         []Alerta            // Relationship with Alertas
	// Accesos         []BitacoraAcceso    // Relationship with Bitacora_Accesos
	// Actividades     []BitacoraActividad // Relationship with Bitacora_Actividades
}

// type Usuario struct {
// 	gorm.Model

// 	ID_usuario      int          `gorm:"not null;unique_index"`      // Standard field for the primary key
// 	Nombre          string       `gorm:"type:varchar(100);not null"` // A regular string field
// 	Apellido        string       `gorm:"type:varchar(100);not null"`
// 	Contrasenia     *string      // A pointer to a string, allowing for null values
// 	Age             uint8        // An unsigned 8-bit integer
// 	Birthday        *time.Time   // A pointer to time.Time, can be null
// 	ActivatedAt     sql.NullTime // Uses sql.NullTime for nullable time fields
// 	Fecha_creaacion time.Time    // Automatically managed by GORM for creation time
// 	Actualizado_el  time.Time
// }
