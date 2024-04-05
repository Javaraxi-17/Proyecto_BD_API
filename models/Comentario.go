package models

import (
	"time"


)


type Comentario struct {
    ID_comentario   uint `gorm:"primaryKey;column:ID_comentario"`
    IDUsuario       uint
    IDCasa          uint
    IDReserva       uint       `gorm:"index"`
    Comentario      string     `gorm:"type:text"`
    Calificacion    int
    FechaPublicacion time.Time
    Usuario         Usuario    `gorm:"foreignKey:IDUsuario"`
    Casa            Casa       `gorm:"foreignKey:IDCasa"`
    Reserva         Reserva    `gorm:"foreignKey:IDReserva"`
}
