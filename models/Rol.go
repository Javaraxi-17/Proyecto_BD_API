package models



type Rol struct {
    ID_rol uint           `gorm:"primaryKey;column:ID_rol"`
    Nombre   string    `gorm:"type:varchar(100);not null"`
    Usuarios []Usuario `gorm:"many2many:usuarios_roles;"`
}