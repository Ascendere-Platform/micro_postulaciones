package apibd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidoUsuario(id string, tk string) (apimodels.UsuarioEquipo, error) {
	var miembro apimodels.UsuarioEquipo

	var usuario struct {
		ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Nombres    string             `bson:"nombre" json:"nombre,omitempty"`
		Apellidos string `bson:"apellidos" json:"apellidos,omitempty"`
		Email      string             `bson:"email" json:"email"`
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://34.123.95.33/verPerfil?id="+id, nil)

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

	json.Unmarshal(bodyBytes, &usuario)

	miembro.Nombres = usuario.Nombres + usuario.Apellidos
	miembro.Email = usuario.Email

	return miembro, err
}
