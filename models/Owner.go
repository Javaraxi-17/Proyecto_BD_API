package models

type Propietario struct {
	IDPropietario     uint   `gorm:"primaryKey;column:id_propietario"`
	Nombre            string `gorm:"column:nombre;type:varchar"`
	CorreoElectronico string `gorm:"column:correo_electronico;type:varchar"`
	Contrasena        string `gorm:"column:contrasena;type:varchar"`
	DetallesContacto  string `gorm:"column:detalles_contacto;type:varchar"`
	// Casas             []Casa `gorm:"foreignKey:IDPropietario"`
}
