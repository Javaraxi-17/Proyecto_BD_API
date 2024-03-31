package models

import (
	"time"

	"gorm.io/gorm"
)

type Comentario struct {
	gorm.Model

	IDUsuario        uint
	IDCasa           uint
	Comentario       string `gorm:"type:text"`
	Calificacion     int
	FechaPublicacion time.Time
}
