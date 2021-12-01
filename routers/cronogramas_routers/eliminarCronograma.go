package cronogramasrouters

import (
	"net/http"

	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
)

func EliminarCronograma(w http.ResponseWriter, r *http.Request) {

	cronograma := r.URL.Query().Get("id")

	if len(cronograma) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := cronogramabd.EliminoCronograma(cronograma)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}