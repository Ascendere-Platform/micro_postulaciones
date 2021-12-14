package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/micro-postulaciones/middlew"
	cronogramasrouters "github.com/ascendere/micro-postulaciones/routers/cronogramas_routers"
	equiporouters "github.com/ascendere/micro-postulaciones/routers/equipo_routers"
	postulacionrouters "github.com/ascendere/micro-postulaciones/routers/postulacion_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Llamadas al CRUD de Cronograma
	router.HandleFunc("/registrarHito", middlew.ChequeoBD(middlew.ValidoJWT(cronogramasrouters.RegistrarCronograma))).Methods("POST")
	router.HandleFunc("/eliminarHito", middlew.ChequeoBD(middlew.ValidoJWT(cronogramasrouters.EliminarCronograma))).Methods("DELETE")
	router.HandleFunc("/buscarHito", middlew.ChequeoBD(middlew.ValidoJWT(cronogramasrouters.BuscarCronograma))).Methods("GET")
	router.HandleFunc("/listarHitos", middlew.ChequeoBD(middlew.ValidoJWT(cronogramasrouters.ListarCronograma))).Methods("GET")

	//Llamadas al CRUD de Postulaciones
	router.HandleFunc("/registrarPostulacion", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.RegistrarPostulacion))).Methods("POST")
	router.HandleFunc("/buscarPostulacion", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.BuscarPostulacion))).Methods("GET")
	router.HandleFunc("/listarPostulaciones", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.ListarPostulaciones))).Methods("GET")
	router.HandleFunc("/eliminarPostulacion", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.EliminarPostulacion))).Methods("DELETE")
	router.HandleFunc("/actualizarPostulacion", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.ActualizarPostulacion))).Methods("PUT")
	router.HandleFunc("/publicarPostulacion", middlew.ChequeoBD(middlew.ValidoJWT(postulacionrouters.PublicarPostulacion))).Methods("PUT")

	//Llamadas al CRUD de Miembros
	router.HandleFunc("/agregarMiembro", middlew.ChequeoBD(middlew.ValidoJWT(equiporouters.AgregarMiembroEquipo))).Methods("POST")
	router.HandleFunc("/eliminarMiembro", middlew.ChequeoBD(middlew.ValidoJWT(equiporouters.EliminoMiembro))).Methods("DELETE")
	router.HandleFunc("/actualizarMiembro", middlew.ChequeoBD(middlew.ValidoJWT(equiporouters.ActualizarMiembroEquipo))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
