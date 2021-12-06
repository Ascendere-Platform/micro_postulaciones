package bd

import (
	"context"
	"time"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParteEquipo(idProyecto primitive.ObjectID, idUsuario primitive.ObjectID) (apimodels.UsuarioEquipo, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Postualciones")
	col := db.Collection("postulacion")

	condicion := bson.M{"_id": idProyecto}

	var usuarioEncontrado apimodels.UsuarioEquipo

	var postulacion postulacionmodels.Postulacion

	err := col.FindOne(ctx, condicion).Decode(postulacion)

	if err != nil {
		return usuarioEncontrado, false, err
	}

	for _, miembro := range postulacion.Equipo{
		if miembro.ID == idUsuario {
			usuarioEncontrado = miembro
			break 
		}
	}

	return usuarioEncontrado, true, nil

}

func ValidoGestor(idProyecto primitive.ObjectID, idUsuario primitive.ObjectID) (apimodels.UsuarioEquipo, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Postualciones")
	col := db.Collection("postulacion")

	condicion := bson.M{"_id": idProyecto}

	var usuarioEncontrado apimodels.UsuarioEquipo

	var postulacion postulacionmodels.Postulacion

	err := col.FindOne(ctx, condicion).Decode(postulacion)

	if err != nil {
		return usuarioEncontrado, false, err
	}

	for _, miembro := range postulacion.Equipo{
		if miembro.ID == idUsuario && miembro.Cargo == "GESTOR" {
			usuarioEncontrado = miembro
			break 
		}
	}

	return usuarioEncontrado, true, nil

}