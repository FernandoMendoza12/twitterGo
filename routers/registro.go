package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FernandoMendoza12/twitterGo/bd"
	"github.com/FernandoMendoza12/twitterGo/models"
)

func Registro(ctx context.Context) models.ResApi {
	var t models.Usuario
	var r models.ResApi
	r.Status = 400

	fmt.Println("Se ejecuto Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "El password debe de contener mas de 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario con ese email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		r.Message = "Ocurrio un error al registrar el usuario" + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se a logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "El registro del usuario se hizo de manera correcta"
	fmt.Println(r.Message)
	return r

}
