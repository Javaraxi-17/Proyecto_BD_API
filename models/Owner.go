package models

import "gorm.io/gorm"

type Propietario struct {
	gorm.Model

	Nombre           string `gorm:"type:varchar(100);not null"`
	Correo           string `gorm:"type:varchar(100);not null"`
	Contrasena       string `gorm:"type:varchar(100);not null"`
	DetallesContacto string `gorm:"type:varchar(255)"`
	// Casas            []Casa                    // Relationship with Casas
	// CasasSeguimiento []Casa                    `gorm:"many2many:seguimiento_casas;"` // Relationship with Casas through Seguimiento_Casas
	// Comisiones       []ComisionAgente          // Relationship with Comisiones_Agentes
	// SolicitudesAdmin []SolicitudAdministrativa // Relationship with Solicitudes_Administrativas
}
