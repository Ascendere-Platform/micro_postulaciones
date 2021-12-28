package apibd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
)

func DevuelvoUsuarioEquipo(usuario apimodels.UsuarioEquipo, tk string) (apimodels.DevuelvoUsuarioEquipo, error) {
	var miembro apimodels.DevuelvoUsuarioEquipo

	miembro.ID = usuario.ID
	miembro.Nombres = usuario.Nombres
	miembro.Email = usuario.Email
	miembro.Cargo = usuario.Cargo

	var asignatura apimodels.Asignatura

	id := usuario.Asignatura.Hex()

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://35.232.65.39/buscarAsignatura?asignatura="+id, nil)

	if err != nil {
		miembro.Asignatura.ID = miembro.ID
		miembro.Asignatura.NombreAsignatura = ""
		miembro.Asignatura.FacultadID = ""
		miembro.Asignatura.Modalidad = ""
		return miembro, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tk)

	resp, error := client.Do(req)

	if error != nil {
		miembro.Asignatura.ID = miembro.ID
		miembro.Asignatura.NombreAsignatura = ""
		miembro.Asignatura.FacultadID = ""
		miembro.Asignatura.Modalidad = ""
		return miembro, error
	}

	defer resp.Body.Close()

	bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

	if errorBytes != nil {
		miembro.Asignatura.ID = miembro.ID
		miembro.Asignatura.NombreAsignatura = ""
		miembro.Asignatura.FacultadID = ""
		miembro.Asignatura.Modalidad = ""
		return miembro, errorBytes
	}

	json.Unmarshal(bodyBytes, &asignatura)

	miembro.Asignatura = asignatura

	return miembro, err
}
