package models

type AgenteInmobiliario struct {
	ID_agente_inmobiliario uint             `gorm:"primaryKey;column:ID_agente_inmobiliario"`
	Nombre                 string           `gorm:"type:varchar(100);not null"`
	Correo                 string           `gorm:"type:varchar(100);not null;unique"`
	Contrasena             string           `gorm:"type:varchar(100);not null"`
	DetallesContacto       string           `gorm:"type:varchar(255)"`
	Casas                  []Casa           `gorm:"foreignKey:IDAgente"`
	Comisiones             []ComisionAgente `gorm:"foreignKey:IDAgente"`
}
