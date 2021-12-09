package cronogramasrouters

import (
	"encoding/json"
	"net/http"

	"github.com/ascendere/micro-postulaciones/bd"
	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	cronogramamodels "github.com/ascendere/micro-postulaciones/models/cronograma_models"
	"github.com/ascendere/micro-postulaciones/routers"
)

func RegistrarCronograma(w http.ResponseWriter, r *http.Request) {
	var cronograma cronogramamodels.Cronograma

	err := json.NewDecoder(r.Body).Decode(&cronograma)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	propuesta, errPropuesta := postulacionbd.BuscoPostulacion(cronograma.PostualcionId.Hex(), routers.Tk)

	if errPropuesta != nil {
		http.Error(w, "El hito no se puede registrar por que el id de la postulación no se encuentra o es incorrecto", 402)
		return
	}

	_, encontrado := bd.ParteEquipo(propuesta, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No es parte del equipo, no puede registrar un nuevo Hito al cronograma", 401)
		return
	}

	_, status, err := cronogramabd.RegistrarCronograma(cronograma)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo Hito del cronograma", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado registrar un nuevo Hito del cronograma", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
