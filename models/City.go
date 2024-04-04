package models



type Ciudad struct {
    ID           uint
    NombreCiudad string `gorm:"type:varchar(100);not null;unique"`
    PaisID       uint
    Pais         Pais   `gorm:"foreignKey:PaisID"`
    Barrios      []Barrio `gorm:"many2many:ciudad_barrio;"`
}

type Barrio struct {
    ID          uint
    NombreBarrio string `gorm:"type:varchar(100);not null;unique"`
    Ciudades     []Ciudad `gorm:"many2many:ciudad_barrio;"`
}