package postulacionrouters

import (
	"net/http"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	"github.com/ascendere/micro-postulaciones/routers"
)

func EliminarPostulacion(w http.ResponseWriter, r *http.Request) {

	postulacion := r.URL.Query().Get("id")

	if len(postulacion) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}
	informacion, _ := postulacionbd.BuscoPostulacion(postulacion, routers.Tk)

	_, encontrado := bd.ValidoGestor(informacion, routers.IDUsuario)

	if !encontrado {
		http.Error(w, "No esta autorizado a eliminar la postulación", 401)
		return
	}

	err := postulacionbd.EliminoPostulacion(postulacion)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}