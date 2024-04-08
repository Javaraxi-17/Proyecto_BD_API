package models

type CategoriaPropiedad struct {
	ID_categoria_propiedad uint   `gorm:"primaryKey;column:ID_categoria_propiedad"`
	Nombre                 string `gorm:"type:varchar(100);not null;unique"`
	Descripcion            string `gorm:"type:varchar(255)"`
	// Casas                  []Casa `gorm:"many2many:casa_categoria;foreignKey:ID_categoria_propiedad;joinForeignKey:ID_categoria_propiedad;References:ID_casa;joinReferences:Id_casa"`
}

type CasaCategoria struct {
	IDCasa                 uint   `gorm:"primaryKey;column:ID_casa"`
	ID_categoria_propiedad uint   `gorm:"primaryKey;column:ID_categoria_propiedad"`
	CasaCategoria_nombre   string `gorm:"type:varchar(255)"`
}
