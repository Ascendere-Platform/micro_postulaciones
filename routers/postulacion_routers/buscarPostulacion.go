package postulacionrouters

import (
	"encoding/json"
	"log"
	"net/http"

	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	"github.com/ascendere/micro-postulaciones/routers"
)

func BuscarPostulacion(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := postulacionbd.BuscoPostulacion(id, routers.Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una postulacion " + err.Error(), 400)
		return
	}

	log.Println(informacion.Estado)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}