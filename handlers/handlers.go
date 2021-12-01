package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/micro-postulaciones/middlew"
	cronogramasrouters "github.com/ascendere/micro-postulaciones/routers/cronogramas_routers"
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


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}