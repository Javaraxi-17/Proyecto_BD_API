package models

import "gorm.io/gorm"

type AgenteInmobiliario struct {
	gorm.Model

	Nombre           string `gorm:"type:varchar(100);not null"`
	Correo           string `gorm:"type:varchar(100);not null"`
	Contrasena       string `gorm:"type:varchar(100);not null"`
	DetallesContacto string `gorm:"type:varchar(255)"`
	// Casas            []Casa                    // Relationship with Casas
	// Comisiones       []ComisionAgente          // Relationship with Comisiones_Agentes
	// SolicitudesAdmin []SolicitudAdministrativa // Relationship with Solicitudes_Administrativas
}
