package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando apenas um usuarios"))
}

func BuscarVariosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando varios usuarios"))
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewRepositoryUser(db)
	userID, erro := repository.Created(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted with success: %d", userID)))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
