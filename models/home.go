package models

import "gorm.io/gorm"

type Casa struct {
	gorm.Model

	IDPropietario       uint
	IDAgente            uint
	Direccion           string `gorm:"type:varchar(255);not null"`
	Ciudad              string `gorm:"type:varchar(100);not null"`
	Pais                string `gorm:"type:varchar(100);not null"`
	Descripcion         string `gorm:"type:varchar(255)"`
	Precio              float64
	TipoPropiedad       string `gorm:"type:varchar(100);not null"`
	NumHabitaciones     int
	NumBanos            int
	Amenidades          string `gorm:"type:text"`
	AreaMetrosCuadrados int
	Fotos               string `gorm:"type:text"`
	// Categorias          []CategoriaPropiedad `gorm:"many2many:casa_categoria;"`
	// Comentarios         []Comentario         // Relationship with Comentarios
	// Reservas            []Reserva            // Relationship with Reservas
	// Seguimientos        []SeguimientoCasa    // Relationship with Seguimiento_Casas
	// HistorialPrecios    []HistorialPrecio    // Relationship with Historial_Precios
}
