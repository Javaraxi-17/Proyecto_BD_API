package main

// Antes de ejecutar main.go, asegúrate de que las relaciones estén comentadas.

// Después de haber ejecutado main.go y creado las tablas, descomenta las relaciones.
// Elimina los comentarios de las relaciones en tus modelos para que las asociaciones funcionen correctamente.

import (
	"log"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/Javaraxi-17/Proyecto_BD_API/routes"
	"github.com/gorilla/mux"
)

func migrateModels() {
	db.DB.AutoMigrate(
		&models.AgenteInmobiliario{},
		&models.Propietario{},
		&models.Usuario{},
		&models.Pais{},
		&models.Ciudad{},
		&models.Barrio{},
		&models.CategoriaPropiedad{},
		&models.Casa{},
		&models.Reserva{},
		&models.Transaccion{},
		&models.Comentario{},
		&models.ReservaHistorica{},
		&models.PagoPendiente{},
		&models.Promocion{},
		&models.AplicacionPromocion{},
		&models.Alerta{},
		&models.BitacoraAcceso{},
		&models.BitacoraActividad{},
		&models.Configuracion{},
		&models.ComisionAgente{},
		&models.SeguimientoCasa{},
		&models.HistorialPrecio{},
		&models.SolicitudAdministrativa{},
		&models.Auditoria{},
		&models.Rol{},
	)
}

func main() {
	// Conectar a la base de datos
	db.DBConn()

	// Migrar modelos
	migrateModels()

	// // Insertar datos de ejemplo

	// models.InsertDataCategoriasCasa()
	// models.InsertarUsuarios()
	// models.InsertarRoles()
	// models.InsertarPaises()
	// models.InsertarCiudades()
	// models.InsertarBarrios()
	// models.InsertarCategoriasPropiedad()

	// Configurar rutas
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//owner
	r.HandleFunc("/owners", routes.GetOwnersHandler).Methods("GET")
	r.HandleFunc("/owners/{id}", routes.GetOwnerHandler).Methods("GET")
	r.HandleFunc("/owners", routes.PostOwnerHandler).Methods("POST")
	r.HandleFunc("/owners/{id}", routes.DeleteOwnerHandler).Methods("DELETE")

	//Home
	r.HandleFunc("/homes", routes.GetHomesHandler).Methods("GET")
	r.HandleFunc("/homes/{id}", routes.GetHomeHandler).Methods("GET")
	r.HandleFunc("/homes", routes.PostHomeHandler).Methods("POST")
	r.HandleFunc("/homes/{id}", routes.DeleteHomesHandler).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
