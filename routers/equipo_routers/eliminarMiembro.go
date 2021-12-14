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

func EliminoMiembro(w http.ResponseWriter, r *http.Request) {
	var miembro apimodels.UsuarioEquipo
	id := r.URL.Query().Get("id")
	err := json.NewDecoder(r.Body).Decode(&miembro)

	if err != nil {
		http.Error(w, "Ocurrio un error al eliminar un miembro del equipo ", http.StatusBadRequest)
		return
	}

	informacion, _ := postulacionbd.BuscoPostulacion(id, routers.Tk)

	_, encontrado := bd.ValidoGestor(informacion, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No esta autorizado a eliminar un miembro del equipo", 401)
		return
	}

	status, errUpdt := equipobd.EliminoMiembro(id, miembro)

	if errUpdt != nil {
		http.Error(w, "Problema al eliminar "+ status + " " +errUpdt.Error(), 402)
		return
	}

	if len(status) != 0 {
		http.Error(w, "Error en: "+status, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}