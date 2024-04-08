package models

type Ciudad struct {
	IDCiudad uint     `gorm:"primaryKey;column:id_ciudad"`
	Nombre   string   `gorm:"column:nombre;type:varchar"`
	IDPais   uint     `gorm:"column:id_pais"` // Clave foránea que referencia a Pais
	Pais     Pais     `gorm:"foreignKey:IDPais"`
	Barrios  []Barrio `gorm:"foreignKey:IDCiudad"`
	Casas    []Casa   `gorm:"foreignKey:IDCiudad"`
}

type Barrio struct {
	ID_barrio    uint   `gorm:"primaryKey;column:ID_barrio"`
	NombreBarrio string `gorm:"type:varchar(100);not null;unique"`
	IDCiudad     uint   `gorm:"column:id_ciudad"`
	Ciudad       Ciudad `gorm:"foreignKey:IDCiudad"`
}
