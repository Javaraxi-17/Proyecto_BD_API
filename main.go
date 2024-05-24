package main

import (
	"log"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/Javaraxi-17/Proyecto_BD_API/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors" // Importa la biblioteca cors
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

	// Insertar datos de ejemplo
	models.InsertarUsuarios()
	models.InsertarRoles()
	models.InsertarPaises()
	models.InsertarCiudades()
	models.InsertarBarrios()
	models.InsertarCategoriasPropiedad()

	// Configurar rutas
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/login", routes.AuthenticateUserHandler).Methods("POST")

	// Owner
	r.HandleFunc("/owners", routes.GetOwnersHandler).Methods("GET")
	r.HandleFunc("/owners/{id}", routes.GetOwnerHandler).Methods("GET")
	r.HandleFunc("/owners", routes.PostOwnerHandler).Methods("POST")
	r.HandleFunc("/owners/{id}", routes.DeleteOwnerHandler).Methods("DELETE")

	// Home
	r.HandleFunc("/homes", routes.GetHomesHandler).Methods("GET")
	r.HandleFunc("/homes/{id}", routes.GetHomeHandler).Methods("GET")
	r.HandleFunc("/homes", routes.PostHomeHandler).Methods("POST")
	r.HandleFunc("/homes/{id}", routes.DeleteHomesHandler).Methods("DELETE")

	// Rol
	r.HandleFunc("/roles", routes.GetRolesHandler).Methods("GET")
	r.HandleFunc("/roles/{id}", routes.GetRolHandler).Methods("GET")
	r.HandleFunc("/roles", routes.PostRolHandler).Methods("POST")
	r.HandleFunc("/roles/{id}", routes.DeleteRolesHandler).Methods("DELETE")

	// Transaction
	r.HandleFunc("/transactions", routes.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", routes.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions", routes.PostTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{id}", routes.DeleteTransactionsHandler).Methods("DELETE")

	// Agente inmobiliario
	r.HandleFunc("/agents", routes.GetAgentesInmobiliariosHandler).Methods("GET")
	r.HandleFunc("/agents/{id}", routes.GetAgenteInmobiliarioHandler).Methods("GET")
	r.HandleFunc("/agents", routes.PostAgenteInmobiliarioHandler).Methods("POST")
	r.HandleFunc("/agents/{id}", routes.DeleteAgentesInmobiliariosHandler).Methods("DELETE")

	// Promocion
	r.HandleFunc("/promos", routes.GetPromocionesHandler).Methods("GET")
	r.HandleFunc("/promos/{id}", routes.GetPromocionHandler).Methods("GET")
	r.HandleFunc("/promos", routes.PostPromocionHandler).Methods("POST")
	r.HandleFunc("/promos/{id}", routes.DeletePromocionesHandler).Methods("DELETE")

	// Aplicacion de promocion
	r.HandleFunc("/apromos", routes.GetAplicacionPromocionesHandler).Methods("GET")
	r.HandleFunc("/apromos/{id}", routes.GetAplicacionPromocionHandler).Methods("GET")
	r.HandleFunc("/apromos", routes.PostAplicacionPromocionHandler).Methods("POST")
	r.HandleFunc("/apromos/{id}", routes.DeleteAplicacionPromocionesHandler).Methods("DELETE")

	// Reserva
	r.HandleFunc("/reserva", routes.GetReservasHandler).Methods("GET")
	r.HandleFunc("/reserva/{id}", routes.GetReservaHandler).Methods("GET")
	r.HandleFunc("/reserva", routes.PostReservaHandler).Methods("POST")
	r.HandleFunc("/reserva/{id}", routes.DeleteReservasHandler).Methods("DELETE")

	// ReservaHistorica
	r.HandleFunc("/reservah", routes.GetReservasHistoricasHandler).Methods("GET")
	r.HandleFunc("/reservah/{id}", routes.GetReservaHistoricaHandler).Methods("GET")
	r.HandleFunc("/reservah", routes.PostReservaHistoricaHandler).Methods("POST")
	r.HandleFunc("/reservah/{id}", routes.DeleteReservasHistoricasHandler).Methods("DELETE")

	// Pago pendiente
	r.HandleFunc("/pagopendiente", routes.GetPagosPendientesHandler).Methods("GET")
	r.HandleFunc("/pagopendiente/{id}", routes.GetPagoPendienteHandler).Methods("GET")
	r.HandleFunc("/pagopendiente", routes.PostPagoPendienteHandler).Methods("POST")
	r.HandleFunc("/pagopendiente/{id}", routes.DeletePagosPendientesHandler).Methods("DELETE")

	// Comision agente
	r.HandleFunc("/comision", routes.GetComisionAgentesHandler).Methods("GET")
	r.HandleFunc("/comision/{id}", routes.GetComisionAgenteHandler).Methods("GET")
	r.HandleFunc("/comision", routes.PostComisionAgenteHandler).Methods("POST")
	r.HandleFunc("/comision/{id}", routes.DeleteComisionAgentesHandler).Methods("DELETE")

	// Alertas
	r.HandleFunc("/alerta", routes.GetAlertasHandler).Methods("GET")
	r.HandleFunc("/alerta/{id}", routes.GetAlertaHandler).Methods("GET")
	r.HandleFunc("/alerta", routes.PostAlertaHandler).Methods("POST")
	r.HandleFunc("/alerta/{id}", routes.DeleteAlertasHandler).Methods("DELETE")

	// Configuracion
	r.HandleFunc("/configuracion", routes.GetConfiguracionesHandler).Methods("GET")
	r.HandleFunc("/configuracion/{id}", routes.GetConfiguracionHandler).Methods("GET")
	r.HandleFunc("/configuracion", routes.PostConfiguracionHandler).Methods("POST")
	r.HandleFunc("/configuracion/{id}", routes.DeleteConfiguracionesHandler).Methods("DELETE")

	// Historial precio
	r.HandleFunc("/historial", routes.GetHistorialPreciosHandler).Methods("GET")
	r.HandleFunc("/historial/{id}", routes.GetHistorialPrecioHandler).Methods("GET")
	r.HandleFunc("/historial", routes.PostHistorialPrecioHandler).Methods("POST")
	r.HandleFunc("/historial/{id}", routes.DeleteHistorialPreciosHandler).Methods("DELETE")

	// Solicitud administrativa
	r.HandleFunc("/solicitud", routes.GetSolicitudesAdministrativasHandler).Methods("GET")
	r.HandleFunc("/solicitud/{id}", routes.GetSolicitudAdministrativaHandler).Methods("GET")
	r.HandleFunc("/solicitud", routes.PostSolicitudAdministrativaHandler).Methods("POST")
	r.HandleFunc("/solicitud/{id}", routes.DeleteSolicitudesAdministrativasHandler).Methods("DELETE")

	// Paises
	r.HandleFunc("/paises", routes.GetPaisesHandler).Methods("GET")
	r.HandleFunc("/paises/{id}", routes.GetPaisHandler).Methods("GET")
	r.HandleFunc("/paises", routes.PostPaisHandler).Methods("POST")
	r.HandleFunc("/paises/{id}", routes.DeletePaisesHandler).Methods("DELETE")

	// Comentarios
	r.HandleFunc("/comentario", routes.GetComentariosHandler).Methods("GET")
	r.HandleFunc("/comentario/{id}", routes.GetComentarioHandler).Methods("GET")
	r.HandleFunc("/comentario", routes.PostComentarioHandler).Methods("POST")
	r.HandleFunc("/comentario/{id}", routes.DeleteComentariosHandler).Methods("DELETE")

	// Ciudad
	r.HandleFunc("/ciudad", routes.GetCiudadesHandler).Methods("GET")
	r.HandleFunc("/ciudad/{id}", routes.GetCiudadHandler).Methods("GET")
	r.HandleFunc("/ciudad", routes.PostCiudadHandler).Methods("POST")
	r.HandleFunc("/ciudad/{id}", routes.DeleteCiudadesHandler).Methods("DELETE")

	// Barrio
	r.HandleFunc("/barrio", routes.GetBarriosHandler).Methods("GET")
	r.HandleFunc("/barrio/{id}", routes.GetBarrioHandler).Methods("GET")
	r.HandleFunc("/barrio", routes.PostBarrioHandler).Methods("POST")
	r.HandleFunc("/barrio/{id}", routes.DeleteBarriosHandler).Methods("DELETE")

	// Categoria propiedad
	r.HandleFunc("/cpropiedad", routes.GetCategoriaPropiedadesHandler).Methods("GET")
	r.HandleFunc("/cpropiedad/{id}", routes.GetCategoriaPropiedadHandler).Methods("GET")
	r.HandleFunc("/cpropiedad", routes.PostCategoriaPropiedadHandler).Methods("POST")
	r.HandleFunc("/cpropiedad/{id}", routes.DeleteCategoriaPropiedadesHandler).Methods("DELETE")

	// Bitacora acceso
	r.HandleFunc("/bitacceso", routes.GetBitacoraAccesosHandler).Methods("GET")
	r.HandleFunc("/bitacceso/{id}", routes.GetBitacoraAccesoHandler).Methods("GET")
	r.HandleFunc("/bitacceso", routes.PostBitacoraAccesoHandler).Methods("POST")
	r.HandleFunc("/bitacceso/{id}", routes.DeleteBitacoraAccesosHandler).Methods("DELETE")

	// Bitacora actividad
	r.HandleFunc("/bitactividad", routes.GetBitacoraActividadesHandler).Methods("GET")
	r.HandleFunc("/bitactividad/{id}", routes.GetBitacoraActividadHandler).Methods("GET")
	r.HandleFunc("/bitactividad", routes.PostBitacoraActividadHandler).Methods("POST")
	r.HandleFunc("/bitactividad/{id}", routes.DeleteBitacoraActividadesHandler).Methods("DELETE")

	// Auditorias
	r.HandleFunc("/auditorias", routes.GetAuditoriasHandler).Methods("GET")
	r.HandleFunc("/auditorias/{id}", routes.GetAuditoriaHandler).Methods("GET")
	r.HandleFunc("/auditorias", routes.PostAuditoriaHandler).Methods("POST")
	r.HandleFunc("/auditorias/{id}", routes.DeleteAuditoriasHandler).Methods("DELETE")

	// Seguimiento de casas
	r.HandleFunc("/segcasas", routes.GetSeguimientoCasasHandler).Methods("GET")
	r.HandleFunc("/segcasas/{id}", routes.GetSeguimientoCasaHandler).Methods("GET")
	r.HandleFunc("/segcasas", routes.PostSeguimientoCasaHandler).Methods("POST")
	r.HandleFunc("/segcasas/{id}", routes.DeleteSeguimientoCasasHandler).Methods("DELETE")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir todos los orígenes, puedes cambiarlo a los orígenes específicos que necesites
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Iniciar el servidor con el middleware de CORS
	handler := c.Handler(r)
	log.Println("Starting server on :3001")
	http.ListenAndServe(":3001", handler)
}
