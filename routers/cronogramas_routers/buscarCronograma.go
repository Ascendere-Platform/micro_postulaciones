package cronogramasrouters

import (
	"encoding/json"
	"net/http"

	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
)

func BuscarCronograma(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := cronogramabd.BuscoCronograma(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un Hito del Cronograma ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}