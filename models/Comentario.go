package models

import (
	"time"


)


type Comentario struct {
    ID_comentario   uint `gorm:"primaryKey;column:ID_comentario"`
    IDUsuario       uint
    IDCasa          uint
    Comentario      string     `gorm:"type:text"`
    Calificacion    int
    FechaPublicacion time.Time

}

