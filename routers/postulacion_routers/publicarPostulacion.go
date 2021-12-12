package postulacionrouters

import (
	"encoding/json"
	"net/http"
	"time"

	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"github.com/ascendere/micro-postulaciones/routers"
)

func PublicarPostulacion(w http.ResponseWriter, r*http.Request) {
	if routers.Rol != "admin" {
		http.Error(w, "No esta autorizado a publicar el proyecto", 401)
		return
	}

	var postulacion postulacionmodels.Postulacion

	err := json.NewDecoder(r.Body).Decode(&postulacion)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(),400)
		return
	}

	postulacion.Mensaje = "Publicado por: " + routers.Nombre
	postulacion.FechaPublicacion = time.Now()

	var status bool
	status, err = postulacionbd.ValidoPostulacion(postulacion)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar publicar la postulacion"+err.Error(),400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado publicar la postulacion",400)
		return
	}

	w.WriteHeader(http.StatusCreated)


}