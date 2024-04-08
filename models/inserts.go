package models

import (
	"github.com/Javaraxi-17/Proyecto_BD_API/db"
)

// func InsertUsers(db *gorm.DB) {
// 	users := []Usuario{
// 		{Name: "Alice", Email: "alice@example.com"},
// 		{Name: "Bob", Email: "bob@example.com"},
// 		{Name: "Charlie", Email: "charlie@example.com"},
// 	}

// 	for _, user := range users {
// 		result := db.Create(&user)
// 		if result.Error != nil {
// 			fmt.Println("Error inserting user:", result.Error)
// 		}
// 	}
// }

func InsertDataCategoriasCasa() {
	// Insertar una CategoriaPropiedad
	categoria := CategoriaPropiedad{
		Nombre:      "Apartamentos",
		Descripcion: "Propiedades tipo apartamento",
	}
	db.DB.Create(&categoria)

	// Insertar una CasaCategoria
	casaCategoria := CasaCategoria{
		IDCasa:                 1, // Supongamos que estas IDs existen
		ID_categoria_propiedad: categoria.ID_categoria_propiedad,
		CasaCategoria_nombre:   "casa 1 campos",
	}
	db.DB.Create(&casaCategoria)
}
