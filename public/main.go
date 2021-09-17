package creusuarios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bufio"
	"strings"
	
	connectiondb"golang/conexion"
	infousuario"golang/infoUsuario"
)

var id int
var users map[int] infousuario

func CrearUsuario() http.HandlerFunc {
	fmt.Println("Usuario creado exitosamente!")
	username := readLine()

	fmt.Print("el puntaje del usuario es: ")
	puntaje := readLine()

	id++
	user := usuario { id, username, puntaje }
	users[id] = infousuario

	if err != nil {
		fmt.Println("ERROR CrearUsuario  username: ", err)
	}

	db := connectiondb.DBConnect()
}

func puntajeUsuario()  {
	for id, user := range users {
		fmt.Println(id, "_", user.username)
	}
}

func main() {
	var reader bufio.Reader
	
	port := ":3333"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})


	fmt.Println("API REST Listen on ", port)
	http.ListenAndServe(port, r)
}