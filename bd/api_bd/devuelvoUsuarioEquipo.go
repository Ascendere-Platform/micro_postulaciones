package apibd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
)

func DevuelvoUsuarioEquipo(usuario apimodels.UsuarioEquipo, tk string) (apimodels.DevuelvoUsuarioEquipo, error) {
	miembro := apimodels.DevuelvoUsuarioEquipo {
		ID: usuario.ID,
		Nombres: usuario.Email,
		Email: usuario.Email,
		Cargo: usuario.Cargo,
	}

	var asignatura apimodels.Asignatura

	id := usuario.Asignatura.Hex()

	client := &http.Client{}


	req, err := http.NewRequest("GET", "http://34.123.95.33?asignatura="+id, nil)

	if err != nil {
		return miembro, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tk)

	resp, error := client.Do(req)

	if error != nil {
		return miembro, error
	}

	defer resp.Body.Close()

	bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

	if errorBytes != nil {
		return miembro, errorBytes
	}

	json.Unmarshal(bodyBytes, &asignatura)

	miembro.Asignatura = asignatura

	return miembro, err
}