package controllers

import "net/http"

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando apenas um usuarios"))
}

func BuscarVariosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando varios usuarios"))
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
