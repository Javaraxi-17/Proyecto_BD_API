package models

type CategoriaPropiedad struct {
	ID_categoria_propiedad uint   `gorm:"primaryKey;column:ID_categoria_propiedad"`
	Nombre                 string `gorm:"type:varchar(100);not null;unique"`
	Descripcion            string `gorm:"type:varchar(255)"`
	Casas                  []Casa `gorm:"many2many:casa_categoria;foreignKey:ID_categoria_propiedad;joinForeignKey:ID_categoria_propiedad;References:ID_casa;joinReferences:ID_casa"`
}
