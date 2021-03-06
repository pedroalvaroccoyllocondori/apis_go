package middleware

import (
	"log"
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/autorizacion"
)

func Log(funcion func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion %q ,metodo %q", r.URL.Path, r.Method)
		funcion(w, r) // ejecutar la funcion
	}
}

// funcion de autentificacion de el token
func Autentificacion(funcion func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := autorizacion.ValidarToken(token)
		if err != nil {
			//devolver una respuesta
			prohibido(w, r)
			return

		}
		funcion(w, r) //ejecutar la funcion
	}
}

func prohibido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("no tiene autorizacion"))

}
