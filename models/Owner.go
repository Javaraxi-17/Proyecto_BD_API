package models



type Propietario struct {
    ID               uint
    Nombre           string `gorm:"type:varchar(100);not null"`
    Correo           string `gorm:"type:varchar(100);not null;unique"`
    Contrasena       string `gorm:"type:varchar(100);not null"`
    DetallesContacto string `gorm:"type:varchar(255)"`
    Casas            []Casa `gorm:"foreignKey:IDPropietario"`
}