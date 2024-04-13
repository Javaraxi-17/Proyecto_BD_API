package models

// import (
// 	"github.com/Javaraxi-17/Proyecto_BD_API/db"

// 	 "time"
// )


// func InsertarUsuarios() {
//     usuarios := []Usuario{
//         {Id_usuario: 1, Nombre: "Juan Perez", Correo: "juan@example.com", Contrasena: "contrasena123", Fecha_registro: time.Now()},
//         {Id_usuario: 2, Nombre: "Maria Lopez", Correo: "maria@example.com", Contrasena: "maria123", Fecha_registro: time.Now()},
//         {Id_usuario: 3, Nombre: "Pedro Ramirez", Correo: "pedro@example.com", Contrasena: "pedro456", Fecha_registro: time.Now()},
//         {Id_usuario: 4, Nombre: "Ana Garcia", Correo: "ana@example.com", Contrasena: "ana789", Fecha_registro: time.Now()},
//         {Id_usuario: 5, Nombre: "Luis Rodriguez", Correo: "luis@example.com", Contrasena: "luis456", Fecha_registro: time.Now()},
//         {Id_usuario: 6, Nombre: "Laura Martinez", Correo: "laura@example.com", Contrasena: "laura789", Fecha_registro: time.Now()},
//         {Id_usuario: 7, Nombre: "Carlos Fernandez", Correo: "carlos@example.com", Contrasena: "carlos123", Fecha_registro: time.Now()},
//         {Id_usuario: 8, Nombre: "Sofia Diaz", Correo: "sofia@example.com", Contrasena: "sofia456", Fecha_registro: time.Now()},
//         {Id_usuario: 9, Nombre: "Jorge Herrera", Correo: "jorge@example.com", Contrasena: "jorge789", Fecha_registro: time.Now()},
//         {Id_usuario: 10, Nombre: "Natalia Gomez", Correo: "natalia@example.com", Contrasena: "natalia123", Fecha_registro: time.Now()},
//         {Id_usuario: 11, Nombre: "Diego Sanchez", Correo: "diego@example.com", Contrasena: "diego456", Fecha_registro: time.Now()},
//         {Id_usuario: 12, Nombre: "Valeria Torres", Correo: "valeria@example.com", Contrasena: "valeria789", Fecha_registro: time.Now()},
//         {Id_usuario: 13, Nombre: "Alejandro Ruiz", Correo: "alejandro@example.com", Contrasena: "alejandro123", Fecha_registro: time.Now()},
//         {Id_usuario: 14, Nombre: "Fernanda Garcia", Correo: "fernanda@example.com", Contrasena: "fernanda456", Fecha_registro: time.Now()},
//         {Id_usuario: 15, Nombre: "Roberto Perez", Correo: "roberto@example.com", Contrasena: "roberto789", Fecha_registro: time.Now()},
//         {Id_usuario: 16, Nombre: "Gabriela Rodriguez", Correo: "gabriela@example.com", Contrasena: "gabriela123", Fecha_registro: time.Now()},
//         {Id_usuario: 17, Nombre: "Jose Martinez", Correo: "jose@example.com", Contrasena: "jose456", Fecha_registro: time.Now()},
//         {Id_usuario: 18, Nombre: "Diana Fernandez", Correo: "diana@example.com", Contrasena: "diana789", Fecha_registro: time.Now()},
//         {Id_usuario: 19, Nombre: "Miguel Lopez", Correo: "miguel@example.com", Contrasena: "miguel123", Fecha_registro: time.Now()},
//         {Id_usuario: 20, Nombre: "Paola Ruiz", Correo: "paola@example.com", Contrasena: "paola456", Fecha_registro: time.Now()},
//         {Id_usuario: 21, Nombre: "Arturo Gomez", Correo: "arturo@example.com", Contrasena: "arturo789", Fecha_registro: time.Now()},
//         {Id_usuario: 22, Nombre: "Marina Torres", Correo: "marina@example.com", Contrasena: "marina123", Fecha_registro: time.Now()},
//         {Id_usuario: 23, Nombre: "Ricardo Perez", Correo: "ricardo@example.com", Contrasena: "ricardo456", Fecha_registro: time.Now()},
//         {Id_usuario: 24, Nombre: "Fabiola Diaz", Correo: "fabiola@example.com", Contrasena: "fabiola789", Fecha_registro: time.Now()},
//         {Id_usuario: 25, Nombre: "Santiago Rodriguez", Correo: "santiago@example.com", Contrasena: "santiago123", Fecha_registro: time.Now()},
//         {Id_usuario: 26, Nombre: "Camila Martinez", Correo: "camila@example.com", Contrasena: "camila456", Fecha_registro: time.Now()},
//         {Id_usuario: 27, Nombre: "Manuel Gomez", Correo: "manuel@example.com", Contrasena: "manuel789", Fecha_registro: time.Now()},
//         {Id_usuario: 28, Nombre: "Lucia Lopez", Correo: "lucia@example.com", Contrasena: "lucia123", Fecha_registro: time.Now()},
//         {Id_usuario: 29, Nombre: "Pablo Hernandez", Correo: "pablo@example.com", Contrasena: "pablo456", Fecha_registro: time.Now()},
//         {Id_usuario: 30, Nombre: "Elena Flores", Correo: "elena@example.com", Contrasena: "elena789", Fecha_registro: time.Now()},
//     }

//     for _, usuario := range usuarios {
//         db.DB.Create(&usuario)
//     }

// }


// func InsertarRoles() {
//     rol1 := Rol{
// 		ID_rol :   1,
//         Nombre: "Administrador",
//     }

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

//     db.DB.Create(&rol1)
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

//     // Antigua Guatemala
//     ciudad2 := Ciudad{
//         IDCiudad: 2,
//         Nombre:   "Antigua Guatemala",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad2)

//     // Quetzaltenango
//     ciudad3 := Ciudad{
//         IDCiudad: 3,
//         Nombre:   "Quetzaltenango",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad3)

//     // Escuintla
//     ciudad4 := Ciudad{
//         IDCiudad: 4,
//         Nombre:   "Escuintla",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad4)

//     // Chimaltenango
//     ciudad5 := Ciudad{
//         IDCiudad: 5,
//         Nombre:   "Chimaltenango",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad5)

//     // Huehuetenango
//     ciudad6 := Ciudad{
//         IDCiudad: 6,
//         Nombre:   "Huehuetenango",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad6)

//     // San Marcos
//     ciudad7 := Ciudad{
//         IDCiudad: 7,
//         Nombre:   "San Marcos",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad7)

//     // Retalhuleu
//     ciudad8 := Ciudad{
//         IDCiudad: 8,
//         Nombre:   "Retalhuleu",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad8)

//     // Petén (Santa Elena y Flores)
//     ciudad9 := Ciudad{
//         IDCiudad: 9,
//         Nombre:   "Petén",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad9)

//     ciudad10 := Ciudad{
//         IDCiudad: 10,
//         Nombre:   "Santa Elena",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad10)

//     ciudad11 := Ciudad{
//         IDCiudad: 11,
//         Nombre:   "Flores",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad11)

//     // Cobán
//     ciudad12 := Ciudad{
//         IDCiudad: 12,
//         Nombre:   "Cobán",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad12)

//     // Zacapa
//     ciudad13 := Ciudad{
//         IDCiudad: 13,
//         Nombre:   "Zacapa",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad13)

//     // Chiquimula
//     ciudad14 := Ciudad{
//         IDCiudad: 14,
//         Nombre:   "Chiquimula",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad14)

//     // Puerto Barrios
//     ciudad15 := Ciudad{
//         IDCiudad: 15,
//         Nombre:   "Puerto Barrios",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad15)

//     // Izabal
//     ciudad16 := Ciudad{
//         IDCiudad: 16,
//         Nombre:   "Izabal",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad16)

//     // Jutiapa
//     ciudad17 := Ciudad{
//         IDCiudad: 17,
//         Nombre:   "Jutiapa",
//         IDPais:   1, // ID de Guatemala en la tabla PAIS
//     }
//     db.DB.Create(&ciudad17)
// }



// func InsertarBarrios() {
//     barrios := []Barrio{
//         {ID_barrio: 1, NombreBarrio: "Zona 1 (Centro Histórico)", IDCiudad: 1},
//         {ID_barrio: 2, NombreBarrio: "Zona 10 (Zona Viva)", IDCiudad: 1},
//         {ID_barrio: 3, NombreBarrio: "La Antigua (Centro Histórico)", IDCiudad: 2},
//         {ID_barrio: 4, NombreBarrio: "Jocotenango", IDCiudad: 2},
//         {ID_barrio: 6, NombreBarrio: "Zona 3 (El Centro)", IDCiudad: 3},
//         {ID_barrio: 7, NombreBarrio: "Colonia El Roble", IDCiudad: 4},
//         {ID_barrio: 8, NombreBarrio: "Colonia El Jocotillo", IDCiudad: 4},
//         {ID_barrio: 10, NombreBarrio: "La Trinidad", IDCiudad: 5},
//         {ID_barrio: 12, NombreBarrio: "Barrio Las Vegas", IDCiudad: 6},
//         {ID_barrio: 14, NombreBarrio: "Barrio La Reformita", IDCiudad: 7},
//         {ID_barrio: 15, NombreBarrio: "Colonia El Rosario", IDCiudad: 8},
//         {ID_barrio: 16, NombreBarrio: "Colonia San José", IDCiudad: 8},
//         {ID_barrio: 17, NombreBarrio: "Barrio El Centro (Flores)", IDCiudad: 9},
//         {ID_barrio: 18, NombreBarrio: "Barrio Las Palmas (Santa Elena)", IDCiudad: 9},
//         {ID_barrio: 20, NombreBarrio: "Barrio Santa Cruz", IDCiudad: 10},
//         {ID_barrio: 22, NombreBarrio: "Barrio El Calvario", IDCiudad: 11},
//         {ID_barrio: 24, NombreBarrio: "Barrio La Esperanza", IDCiudad: 12},
//         {ID_barrio: 26, NombreBarrio: "Barrio El Porvenir", IDCiudad: 13},
//         {ID_barrio: 28, NombreBarrio: "Barrio La Unión", IDCiudad: 14},
//         {ID_barrio: 30, NombreBarrio: "Barrio La Democracia", IDCiudad: 15},
//     }

//     for _, barrio := range barrios {
//         db.DB.Create(&barrio)
//     }
// }




// func InsertarCategoriasPropiedad() {
//     categorias := []CategoriaPropiedad{
//         {ID_categoria_propiedad: 1, Nombre: "Casas", Descripcion: "Propiedades residenciales unifamiliares."},
//         {ID_categoria_propiedad: 2, Nombre: "Apartamentos", Descripcion: "Unidades de vivienda en edificios compartidos."},
//         {ID_categoria_propiedad: 3, Nombre: "Chalets", Descripcion: "Casas de estilo alpino con tejado inclinado y balcones."},
//         {ID_categoria_propiedad: 4, Nombre: "Condominios", Descripcion: "Unidades de vivienda en un complejo con servicios compartidos."},
//         {ID_categoria_propiedad: 5, Nombre: "Townhouses", Descripcion: "Casas adosadas en filas o grupos."},
//         {ID_categoria_propiedad: 6, Nombre: "Locales comerciales", Descripcion: "Espacios para negocios y actividades comerciales."},
//         {ID_categoria_propiedad: 7, Nombre: "Oficinas", Descripcion: "Espacios para negocios y actividades administrativas."},
//         {ID_categoria_propiedad: 8, Nombre: "Edificios comerciales", Descripcion: "Edificios destinados principalmente a actividades comerciales."},
//         {ID_categoria_propiedad: 9, Nombre: "Terrenos para desarrollo comercial", Descripcion: "Terrenos destinados a proyectos de desarrollo comercial."},
//         {ID_categoria_propiedad: 10, Nombre: "Bodegas", Descripcion: "Espacios de almacenamiento para mercancías y materiales."},
//         {ID_categoria_propiedad: 11, Nombre: "Naves industriales", Descripcion: "Edificios destinados a actividades industriales."},
//         {ID_categoria_propiedad: 12, Nombre: "Terrenos industriales", Descripcion: "Terrenos destinados a proyectos industriales."},
//         {ID_categoria_propiedad: 13, Nombre: "Terrenos residenciales", Descripcion: "Terrenos destinados a la construcción de viviendas."},
//         {ID_categoria_propiedad: 14, Nombre: "Terrenos comerciales", Descripcion: "Terrenos destinados a actividades comerciales."},
//         {ID_categoria_propiedad: 15, Nombre: "Terrenos agrícolas", Descripcion: "Terrenos destinados a la agricultura o ganadería."},
//         {ID_categoria_propiedad: 16, Nombre: "Casas de playa", Descripcion: "Propiedades residenciales ubicadas en zonas costeras."},
//         {ID_categoria_propiedad: 17, Nombre: "Cabañas de montaña", Descripcion: "Casas rústicas ubicadas en áreas montañosas."},
//         {ID_categoria_propiedad: 18, Nombre: "Casas de campo", Descripcion: "Propiedades residenciales en entornos rurales."},
//         {ID_categoria_propiedad: 19, Nombre: "Villas vacacionales", Descripcion: "Residencias de lujo destinadas al alquiler vacacional."},
//         {ID_categoria_propiedad: 20, Nombre: "Propiedades para renta", Descripcion: "Propiedades destinadas a la renta a largo plazo."},
//         {ID_categoria_propiedad: 21, Nombre: "Propiedades para flip", Descripcion: "Propiedades compradas con el objetivo de renovar y vender rápidamente."},
//         {ID_categoria_propiedad: 22, Nombre: "Desarrollos inmobiliarios", Descripcion: "Proyectos de construcción de múltiples propiedades."},
//         {ID_categoria_propiedad: 23, Nombre: "Propiedades comerciales con retorno de inversión (ROI)", Descripcion: "Propiedades comerciales con un alto potencial de retorno de inversión."},
//     }

//     for _, categoria := range categorias {
//         db.DB.Create(&categoria)
//     }
// }
