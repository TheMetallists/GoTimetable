package main

import (
	"dislexiad/api"
	"dislexiad/api/middleware"
	"dislexiad/database"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Dislexia v. 0.02")

	fmt.Println("Preparing the database...")
	dh := database.DatabaseInit()

	bindTo := "0.0.0.0:9580"
	fmt.Println("Listening on: ", bindTo)

	r := mux.NewRouter()
	// API root
	apir := r.PathPrefix("/api").Subrouter()
	apio := api.ApiObject{
		DH: dh,
	}
	apio.ApiInitRoutes(apir)

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(dh.GetJwtTokenKey()), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	apiAdminRouter := mux.NewRouter()

	n := negroni.New(
		middleware.NewOrigin("*"),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(apiAdminRouter))
	n.Use(negroni.NewLogger())
	apir.PathPrefix("/admin").Handler(n)

	aapir := apiAdminRouter.PathPrefix("/api/admin").Subrouter()

	apio.ApiInitAdminca(aapir)
	aapir.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ASDASDASD!"))
	})

	// temporary: serve unpacked files
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir + "/frontend/build")))

	srv := &http.Server{
		Handler: r,
		Addr:    bindTo,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
