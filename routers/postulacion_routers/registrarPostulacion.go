package postulacionrouters

import (
	"encoding/json"
	"net/http"
	"time"

	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"github.com/ascendere/micro-postulaciones/routers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistrarPostulacion(w http.ResponseWriter, r *http.Request) {
	var postulacion postulacionmodels.Postulacion
	err := json.NewDecoder(r.Body).Decode(&postulacion)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	postulacion.Estado = false
	postulacion.FechaCreacion = time.Now()
	postulacion.Mensaje = "RECIEN CREADO POR : " + routers.Nombre
	postulacion.CalificacionTotal = 0

	objID, _ := primitive.ObjectIDFromHex(routers.IDUsuario)

	gestor := apimodels.UsuarioEquipo{
		ID:    objID,
		Cargo: "GESTOR",
	}

	postulacion.Equipo = append(postulacion.Equipo, gestor)

	_, status, err := postulacionbd.RegistrarPostulacion(postulacion, routers.Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la postulacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la postulación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postulacion)
}
