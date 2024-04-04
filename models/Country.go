package models



type Pais struct {
    ID         uint
    NombrePais string `gorm:"type:varchar(100);not null;unique"`
    Ciudades   []Ciudad `gorm:"foreignKey:PaisID"`
}