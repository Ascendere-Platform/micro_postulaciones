package postulacionrouters

import (
	"encoding/json"
	"net/http"

	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	"github.com/ascendere/micro-postulaciones/routers"
)

func BuscarPostulacion(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := postulacionbd.BuscoPostulacion(id, routers.Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un Hito del Cronograma ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}