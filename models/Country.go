package models



type Pais struct {
    ID_pais    uint `gorm:"primaryKey;column:ID_pais"`
    NombrePais string `gorm:"type:varchar(100);not null;unique"`
    Ciudades   []Ciudad `gorm:"foreignKey:PaisID"`
}