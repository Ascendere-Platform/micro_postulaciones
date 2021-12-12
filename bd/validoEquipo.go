package bd

import (

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParteEquipo(propuesta postulacionmodels.DevuelvoPostulacion, idUsuario string) (apimodels.DevuelvoUsuarioEquipo, bool){
	

	var usuarioEncontrado apimodels.DevuelvoUsuarioEquipo

	var flag bool


	objID,_ := primitive.ObjectIDFromHex(idUsuario)
	for _, miembro := range propuesta.Equipo{
		if miembro.ID == objID {
			usuarioEncontrado = miembro
			flag = true
		}
	}

	return usuarioEncontrado, flag

}

func ValidoGestor(propuesta postulacionmodels.DevuelvoPostulacion, idUsuario string) (apimodels.DevuelvoUsuarioEquipo, bool){

	var usuarioEncontrado apimodels.DevuelvoUsuarioEquipo

	objID,_ := primitive.ObjectIDFromHex(idUsuario)

	var flag bool

	for _, miembro := range propuesta.Equipo{
		if miembro.ID == objID && miembro.Cargo == "GESTOR" {
			usuarioEncontrado = miembro
			flag = true
		}
	}

	return usuarioEncontrado, flag

}