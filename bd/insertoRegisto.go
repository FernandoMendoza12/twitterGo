package bd

import (
	"context"

	"github.com/FernandoMendoza12/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	resultado, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil

}
