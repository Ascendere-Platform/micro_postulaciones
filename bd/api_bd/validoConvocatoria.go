package apibd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
)

func ValidoConvocatoria(id string, tk string) (apimodels.Convocatoria, error) {
	var convocatoria apimodels.Convocatoria

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://34.135.35.238/buscarConvocatoria?id="+id, nil)

	if err != nil {
		return convocatoria, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tk)

	resp, error := client.Do(req)

	if error != nil {
		return convocatoria, error
	}

	defer resp.Body.Close()

	bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

	if errorBytes != nil {
		return convocatoria, errorBytes
	}

	json.Unmarshal(bodyBytes, &convocatoria)

	return convocatoria, err
}

func ValidoTipo (id string, tk string) (apimodels.TipoProyecto, error) {
	var tipoProyecto apimodels.TipoProyecto

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://34.123.204.176/buscarTipoProyecto?id="+id, nil)

	if err != nil {
		return tipoProyecto, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tk)

	resp, error := client.Do(req)

	if error != nil {
		return tipoProyecto, error
	}

	defer resp.Body.Close()

	bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

	if errorBytes != nil {
		return tipoProyecto, errorBytes
	}

	json.Unmarshal(bodyBytes, &tipoProyecto)

	return tipoProyecto, err
}	