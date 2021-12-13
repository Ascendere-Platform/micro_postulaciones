package equiporouters

import (
	"encoding/json"
	"net/http"

	"github.com/ascendere/micro-postulaciones/bd"
	equipobd "github.com/ascendere/micro-postulaciones/bd/equipo_bd"
	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	"github.com/ascendere/micro-postulaciones/routers"
)

func AgregarMiembroEquipo(w http.ResponseWriter, r *http.Request) {
	var miembroNuevo apimodels.UsuarioEquipo
	id := r.URL.Query().Get("id")
	err := json.NewDecoder(r.Body).Decode(&miembroNuevo)

	if err != nil {
		http.Error(w, "Ocurrio un error al agregar un miembro al equipo ", http.StatusBadRequest)
		return
	}

	informacion, _ := postulacionbd.BuscoPostulacion(id, routers.Tk)

	_, encontrado := bd.ValidoGestor(informacion, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No esta autorizado para agregar un nuevo miembro al equipo", 401)
		return
	}

	status, errUpdt := equipobd.AgregoMiembro(id, miembroNuevo)

	if errUpdt != nil {
		http.Error(w, "Problema al actualizar "+ status + " " +errUpdt.Error(), 402)
		return
	}

	if len(status) == 0 {
		http.Error(w, "Error en: "+status, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
