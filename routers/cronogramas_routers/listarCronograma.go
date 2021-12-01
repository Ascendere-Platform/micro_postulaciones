package cronogramasrouters

import (
	"encoding/json"
	"net/http"

	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
)

func ListarCronograma(w http.ResponseWriter, r *http.Request) {

	result, status := cronogramabd.ListoCronograma()
	if !status {
		http.Error(w, "Error al leer los anexos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}