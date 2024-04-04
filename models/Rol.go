package models



type Rol struct {
    ID uint           `gorm:"primaryKey;autoIncrement"`
    Nombre   string    `gorm:"type:varchar(100);not null"`
    Usuarios []Usuario `gorm:"many2many:usuarios_roles;"`
}