package main

// Antes de ejecutar main.go, asegúrate de que las relaciones estén comentadas.

// Después de haber ejecutado main.go y creado las tablas, descomenta las relaciones.
	// Elimina los comentarios de las relaciones en tus modelos para que las asociaciones funcionen correctamente.



import (
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/Javaraxi-17/Proyecto_BD_API/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConn()

	db.DB.AutoMigrate(models.Pais{})
	db.DB.AutoMigrate(models.Ciudad{})
	db.DB.AutoMigrate(models.Barrio{})
	db.DB.AutoMigrate(models.CategoriaPropiedad{})
	db.DB.AutoMigrate(models.Casa{})
	db.DB.AutoMigrate(models.CasaCategoria{})
	db.DB.AutoMigrate(models.Reserva{})
	db.DB.AutoMigrate(models.Transaccion{})
	db.DB.AutoMigrate(models.Comentario{})
	db.DB.AutoMigrate(models.ReservaHistorica{})
	db.DB.AutoMigrate(models.PagoPendiente{})
	db.DB.AutoMigrate(models.Promocion{})
	db.DB.AutoMigrate(models.AplicacionPromocion{})
	db.DB.AutoMigrate(models.Alerta{})
	db.DB.AutoMigrate(models.BitacoraAcceso{})
	db.DB.AutoMigrate(models.BitacoraActividad{})
	db.DB.AutoMigrate(models.Configuracion{})
	db.DB.AutoMigrate(models.ComisionAgente{})
	db.DB.AutoMigrate(models.SeguimientoCasa{})
	db.DB.AutoMigrate(models.HistorialPrecio{})
	db.DB.AutoMigrate(models.SolicitudAdministrativa{})
	db.DB.AutoMigrate(models.Auditoria{})
	db.DB.AutoMigrate(models.Usuario{})
	db.DB.AutoMigrate(models.UsuarioRol{})
	db.DB.AutoMigrate(models.Rol{})
	db.DB.AutoMigrate(models.Propietario{})
	db.DB.AutoMigrate(models.AgenteInmobiliario{})





	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
