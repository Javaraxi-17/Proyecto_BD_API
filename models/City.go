package models



type Ciudad struct {
    ID_ciudad    uint `gorm:"primaryKey;column:ID_ciudad"`
    NombreCiudad string `gorm:"type:varchar(100);not null;unique"`
    PaisID       uint
    Pais         Pais   `gorm:"foreignKey:PaisID"`
    Barrios      []Barrio `gorm:"many2many:ciudad_barrio;"`
}

type Barrio struct {
    ID_barrio    uint `gorm:"primaryKey;column:ID_barrio"`
    NombreBarrio string `gorm:"type:varchar(100);not null;unique"`
    Ciudades     []Ciudad `gorm:"many2many:ciudad_barrio;"`
}