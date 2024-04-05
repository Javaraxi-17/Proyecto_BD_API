package models


type Casa struct {
    ID_casa            uint `gorm:"primaryKey;column:ID_casa"`
    IDPropietario      uint
    IDAgente           uint
    Direccion          string    `gorm:"type:varchar(255);not null"`
    Ciudad             string    `gorm:"type:varchar(100);not null"`
    Pais               string    `gorm:"type:varchar(100);not null"`
    Descripcion        string    `gorm:"type:varchar(255)"`
    Precio             float64
    TipoPropiedad      string    `gorm:"type:varchar(100);not null"`
    NumHabitaciones    int
    NumBanos           int
    Amenidades         string    `gorm:"type:text"`
    AreaMetrosCuadrados int
    Fotos              string    `gorm:"type:text"`
    Propietario        Propietario          `gorm:"foreignKey:IDPropietario"`
    AgenteInmobiliario AgenteInmobiliario   `gorm:"foreignKey:IDAgente"`
    Reservas           []Reserva            `gorm:"foreignKey:IDCasa"`
    Comentarios        []Comentario         `gorm:"foreignKey:IDCasa"`
}
