package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Bienvenido a Spider Quotes"))
}