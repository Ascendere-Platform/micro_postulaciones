package cronogramasrouters

import (
	"encoding/json"
	"net/http"

	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
	cronogramamodels "github.com/ascendere/micro-postulaciones/models/cronograma_models"
)

func RegistrarCronograma(w http.ResponseWriter, r *http.Request) {
	var cronograma cronogramamodels.Cronograma

	err := json.NewDecoder(r.Body).Decode(&cronograma)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := cronogramabd.RegistrarCronograma(cronograma)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo Hito del cronograma", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo Hito del cronograma", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}