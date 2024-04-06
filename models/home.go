package models

import (
	"time"


)



type Casa struct {
    ID_casa            uint `gorm:"primaryKey;column:ID_casa"`
    IDPropietario        uint      `gorm:"column:id_propietario"` // Clave foránea que referencia a Propietario
    IDAgente             uint      `gorm:"column:id_agente"`      // Clave foránea que referencia a AgenteInmobiliario
    Direccion          string    `gorm:"type:varchar(255);not null"`
    IDCiudad             uint      `gorm:"column:id_ciudad"`      // Clave foránea que referencia a Ciudad
    Descripcion        string    `gorm:"type:varchar(255)"`
    Precio             float64
    TipoPropiedad      string    `gorm:"type:varchar(100);not null"`
    NumHabitaciones    int
    NumBanos           int
    Amenidades         string    `gorm:"type:text"`
    AreaMetrosCuadrados int
    Fotos              string    `gorm:"type:text"`

    // Comentarios []Comentario `gorm:"foreignKey:IDCasa"`
    // HistorialPrecios []HistorialPrecio `gorm:"foreignKey:IDCasa"`
    

    // Propietario        Propietario        `gorm:"foreignKey:IDPropietario"`
    // AgenteInmobiliario AgenteInmobiliario `gorm:"foreignKey:IDAgente"`
    // Ciudad             Ciudad             `gorm:"foreignKey:IDCiudad"`

    // Categorias []CategoriaPropiedad `gorm:"many2many:casa_categoria;foreignKey:ID_casa;joinForeignKey:ID_casa;References:ID_categoria_propiedad;joinReferences:ID_categoria_propiedad"`
    // Reservas []Reserva `gorm:"foreignKey:IDCasa"`
    // ReservasHistoricas []ReservaHistorica `gorm:"foreignKey:IDCasa"`
    // SeguimientosCasa []SeguimientoCasa `gorm:"foreignKey:IDCasa"`
}


type SeguimientoCasa struct {
	ID_seguimiento_casa uint `gorm:"primaryKey;column:ID_seguimiento_casa"`
	IDUsuario           uint
	IDCasa              uint
	FechaHoraVisita     time.Time
	InteresExpresado    string  `gorm:"type:text"`
    Usuario Usuario `gorm:"foreignKey:IDUsuario"`
    Casa    Casa    `gorm:"foreignKey:IDCasa"`
}