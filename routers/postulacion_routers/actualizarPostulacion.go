package postulacionrouters

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"github.com/ascendere/micro-postulaciones/routers"
)

func ActualizarPostulacion(w http.ResponseWriter, r *http.Request) {

	var postulacion postulacionmodels.Postulacion

	err := json.NewDecoder(r.Body).Decode(&postulacion)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	propuesta, errPropuesta := postulacionbd.BuscoPostulacion(postulacion.ID.Hex(), routers.Tk)

	if errPropuesta != nil {
		http.Error(w, "El id de la postulación no se encuentra o es incorrecto", 402)
		return
	}

	_, encontrado := bd.ParteEquipo(propuesta, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No es parte del equipo, no puede actualizar la propuesta", 401)
		return
	}

	postulacion.Mensaje = "Actualizado por: " + routers.Nombre
	postulacion.FechaActualizacion = time.Now()

	var status bool
	status, err = postulacionbd.ActualizoPostulacion(postulacion)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar actualizar la postulacion"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado actualizar la postulacion", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}