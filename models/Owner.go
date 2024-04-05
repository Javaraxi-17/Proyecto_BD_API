package models



type Propietario struct {
    ID_propietario   uint `gorm:"primaryKey;column:ID_propietario"`
    Nombre           string `gorm:"type:varchar(100);not null"`
    Correo           string `gorm:"type:varchar(100);not null;unique"`
    Contrasena       string `gorm:"type:varchar(100);not null"`
    DetallesContacto string `gorm:"type:varchar(255)"`
    Casas            []Casa `gorm:"foreignKey:IDPropietario"`
}