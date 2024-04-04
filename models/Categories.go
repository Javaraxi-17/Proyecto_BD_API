package models



type CategoriaPropiedad struct {
    ID          uint
    Nombre      string `gorm:"type:varchar(100);not null;unique"`
    Descripcion string `gorm:"type:varchar(255)"`
    Casas       []CasaCategoria `gorm:"foreignKey:IDCategoria"`
}

type CasaCategoria struct {
    ID           uint
    IDCasa       uint
    IDCategoria  uint
    Casa         Casa            `gorm:"foreignKey:IDCasa"`
    Categoria    CategoriaPropiedad `gorm:"foreignKey:IDCategoria"`
}