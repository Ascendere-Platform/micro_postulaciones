package postulacionrouters

import (
	"encoding/json"
	"net/http"

	postulacionbd "github.com/ascendere/micro-postulaciones/bd/postulacion_bd"
	"github.com/ascendere/micro-postulaciones/routers"
)

func ListarPostulaciones(w http.ResponseWriter, r*http.Request){
	typePostulacion := r.URL.Query().Get("tipo")
	search := r.URL.Query().Get("busqueda")

	result, error := postulacionbd.ListoPostulaciones(routers.IDUsuario, routers.Tk, search, typePostulacion)
	if error != nil {
		http.Error(w, "Error al leer las postulaciones"+ error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
