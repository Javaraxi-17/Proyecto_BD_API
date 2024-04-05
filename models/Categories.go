package models

type CategoriaPropiedad struct {
	ID_categoria_propiedad uint            `gorm:"primaryKey;column:ID_categoria_propiedad"`
	Nombre                 string          `gorm:"type:varchar(100);not null;unique"`
	Descripcion            string          `gorm:"type:varchar(255)"`
	Casas                  []CasaCategoria `gorm:"foreignKey:IDCategoria"`
}

type CasaCategoria struct {
	ID_casa_categoria uint `gorm:"primaryKey;column:ID_casa_categoria"`
	IDCasa            uint
	IDCategoria       uint
	Casa              Casa               `gorm:"foreignKey:IDCasa"`
	Categoria         CategoriaPropiedad `gorm:"foreignKey:IDCategoria"`
}
