package cronogramasrouters

import (
	"net/http"

	"github.com/ascendere/micro-postulaciones/bd"
	cronogramabd "github.com/ascendere/micro-postulaciones/bd/cronograma_bd"
	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	"github.com/ascendere/micro-postulaciones/routers"
)

func EliminarCronograma(w http.ResponseWriter, r *http.Request) {

	cronograma := r.URL.Query().Get("id")

	if len(cronograma) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}
	informacion, _ := cronogramabd.BuscoCronograma(cronograma)

	propuesta, errPropuesta := postulacionbd.BuscoPostulacion(informacion.PostualcionId.Hex(), routers.Tk)

	if errPropuesta != nil {
		http.Error(w, "El hito no se puede registrar por que el id de la postulación no se encuentra o es incorrecto", 402)
		return
	}

	_, encontrado := bd.ParteEquipo(propuesta, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No es parte del equipo, no puede registrar un nuevo Hito al cronograma", 401)
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
