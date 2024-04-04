package models



type Task struct {
    Title       string `gorm:"not null;unique"`
    Description string
    Done        bool   `gorm:"default:false"`
    UserID      uint   `gorm:"index"`
    Usuario     Usuario `gorm:"foreignKey:UserID"`
}