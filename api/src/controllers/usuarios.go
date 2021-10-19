package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userID"], 10, 64)

	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	repo := repositories.NewRepositoryUser(db)

	user, erro := repo.SearchId(userId)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	responses.JSON(w, http.StatusOK, user)
}

func BuscarVariosUsuarios(w http.ResponseWriter, r *http.Request) {
	NickOrName := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	repository := repositories.NewRepositoryUser(db)

	users, erro := repository.Search(NickOrName)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	responses.JSON(w, http.StatusOK, users)
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("create"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewRepositoryUser(db)
	user.ID, erro = repository.Created(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	bodyReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var user models.User

	if erro := json.Unmarshal(bodyReq, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	if erro = user.Prepare("update"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewRepositoryUser(db)

	if erro = repository.Update(userId, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewRepositoryUser(db)

	if erro = repository.Delete(userId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
