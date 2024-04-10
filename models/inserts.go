package models

// import (
// 	"github.com/Javaraxi-17/Proyecto_BD_API/db"

// 	// "time"
// )


// despues de la primera ejecucion, puedes descomentar los inserts



// func InsertDataCategoriasCasa() {
// 	// Insertar una CategoriaPropiedad
// 	categoria := CategoriaPropiedad{
// 		Nombre:      "Apartamentos",
// 		Descripcion: "Propiedades tipo apartamento",
// 	}
// 	db.DB.Create(&categoria)

// 	// Insertar una CasaCategoria
// 	casaCategoria := CasaCategoria{
// 		IDCasa:                 1, // Supongamos que estas IDs existen
// 		ID_categoria_propiedad: categoria.ID_categoria_propiedad,
// 		CasaCategoria_nombre:   "casa 1 campos",
// 	}
// 	db.DB.Create(&casaCategoria)
// }




// func InsertarUsuarios() {
//     usuario1 := Usuario{
// 		Id_usuario:       1,
//         Nombre:           "Juan Perez",
// //         Correo: "juan@example.com",
// //         Contrasena:       "contrasena123",
// //         Fecha_registro:    time.Now(),
// //     }

// //     usuario2 := Usuario{
// // 		Id_usuario:       2,
// //         Nombre:           "Maria Lopez",
// //         Correo: "maria@example.com",
// //         Contrasena:       "maria123",
// //         Fecha_registro:    time.Now(),
// //     }

// //     usuario3 := Usuario{
// // 		Id_usuario:       3,
// //         Nombre:           "Pedro Ramirez",
// //         Correo: "pedro@example.com",
// //         Contrasena:       "pedro456",
// //         Fecha_registro:    time.Now(),
// //     }

// //     db.DB.Create(&usuario1)
// //     db.DB.Create(&usuario2)
// //     db.DB.Create(&usuario3)
// // }


// func InsertarRoles() {
//     // rol1 := Rol{
// 	// 	ID_rol :   1,
//     //     Nombre: "Administrador",
//     // }

// 	rol2 := Rol{
// 		ID_rol :   2,
//         Nombre: "Usuario Normal",
//     }

// 	rol3 := Rol{
// 		ID_rol :   3,
//         Nombre: "Anfitrión",
//     }

// 	rol4 := Rol{
// 		ID_rol :   4,
//         Nombre: "Superusuario",
//     }

// 	rol5 := Rol{
// 		ID_rol :   5,
//         Nombre: "Verificador de Propiedades",
//     }

// 	rol6 := Rol{
// 		ID_rol :   6,
//         Nombre: "Soporte al Cliente",
//     }

//     // db.DB.Create(&rol1)
//     db.DB.Create(&rol2)
// 	db.DB.Create(&rol3)
// 	db.DB.Create(&rol4)
// 	db.DB.Create(&rol5)
// 	db.DB.Create(&rol6)
// }


// func InsertarPaises() {
//     pais1 := Pais{
//         ID_pais:  1,
//         NombrePais: "Guatemala ",
//     }

//     pais2 := Pais{
//         ID_pais:  2,
//         NombrePais: "Costa Rica",
//     }

//     pais3 := Pais{
//         ID_pais:  3,
//         NombrePais: "El Salvador",
//     }

//     pais4 := Pais{
//         ID_pais:  4,
//         NombrePais: "Nicaragua",
//     }

//     pais5 := Pais{
//         ID_pais:  5,
//         NombrePais: "Panama",
//     }

//     db.DB.Create(&pais1)
//     db.DB.Create(&pais2)
//     db.DB.Create(&pais3)
//     db.DB.Create(&pais4)
//     db.DB.Create(&pais5)
// }



// func InsertarCiudades() {
//     // Ciudad de Guatemala (Capital)
//     ciudad1 := Ciudad{
//         IDCiudad: 1,
//         Nombre:   "Ciudad de Guatemala",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad1)

//     // // Antigua Guatemala
//     // ciudad2 := Ciudad{
//     //     IDCiudad: 2,
//     //     Nombre:   "Antigua Guatemala",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad2)

//     // // Quetzaltenango
//     // ciudad3 := Ciudad{
//     //     IDCiudad: 3,
//     //     Nombre:   "Quetzaltenango",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad3)

//     // // Escuintla
//     // ciudad4 := Ciudad{
//     //     IDCiudad: 4,
//     //     Nombre:   "Escuintla",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad4)

//     // // Chimaltenango
//     // ciudad5 := Ciudad{
//     //     IDCiudad: 5,
//     //     Nombre:   "Chimaltenango",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad5)

//     // // Huehuetenango
//     // ciudad6 := Ciudad{
//     //     IDCiudad: 6,
//     //     Nombre:   "Huehuetenango",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad6)

//     // // San Marcos
//     // ciudad7 := Ciudad{
//     //     IDCiudad: 7,
//     //     Nombre:   "San Marcos",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad7)

//     // // Retalhuleu
//     // ciudad8 := Ciudad{
//     //     IDCiudad: 8,
//     //     Nombre:   "Retalhuleu",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad8)

//     // // Petén (Santa Elena y Flores)
//     // ciudad9 := Ciudad{
//     //     IDCiudad: 9,
//     //     Nombre:   "Petén",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad9)

//     // ciudad10 := Ciudad{
//     //     IDCiudad: 10,
//     //     Nombre:   "Santa Elena",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad10)

//     // ciudad11 := Ciudad{
//     //     IDCiudad: 11,
//     //     Nombre:   "Flores",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad11)

//     // // Cobán
//     // ciudad12 := Ciudad{
//     //     IDCiudad: 12,
//     //     Nombre:   "Cobán",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad12)

//     // // Zacapa
//     // ciudad13 := Ciudad{
//     //     IDCiudad: 13,
//     //     Nombre:   "Zacapa",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad13)

//     // // Chiquimula
//     // ciudad14 := Ciudad{
//     //     IDCiudad: 14,
//     //     Nombre:   "Chiquimula",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad14)

//     // // Puerto Barrios
//     // ciudad15 := Ciudad{
//     //     IDCiudad: 15,
//     //     Nombre:   "Puerto Barrios",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad15)

//     // // Izabal
//     // ciudad16 := Ciudad{
//     //     IDCiudad: 16,
//     //     Nombre:   "Izabal",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad16)

//     // // Jutiapa
//     // ciudad17 := Ciudad{
//     //     IDCiudad: 17,
//     //     Nombre:   "Jutiapa",
//     //     IDPais:   1, // ID de Guatemala en la tabla PAIS
//     // }
//     // db.DB.Create(&ciudad17)
// }


// func InsertarBarrios() {
//     barrio1 := Barrio{
//         ID_barrio: 1,
//         NombreBarrio:   "Zona 1",
//         IDCiudad: 1, // ID de Ciudad de Guatemala en la tabla CIUDAD
//     }
//     db.DB.Create(&barrio1)

//     // barrio2 := Barrio{
//     //     ID_barrio: 2,
//     //     NombreBarrio:   "Zona 10",
//     //     IDCiudad: 1, // ID de Ciudad de Guatemala en la tabla CIUDAD
//     // }
//     // db.DB.Create(&barrio2)

//     // barrio3 := Barrio{
//     //     ID_barrio: 3,
//     //     NombreBarrio:   "Zona 14",
//     //     IDCiudad: 1, // ID de Ciudad de Guatemala en la tabla CIUDAD
//     // }
//     // db.DB.Create(&barrio3)

//     // Puedes continuar agregando más barrios con sus respectivos IDs de ciudad
// }
