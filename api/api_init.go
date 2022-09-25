package api

import (
	"dislexiad/database"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiObject struct {
	DH *database.DatabaseHolder
}

func (thiz *ApiObject) ApiInitRoutes(apir *mux.Router) {
	apir.HandleFunc("/telek", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerTelek(writer, request)
	})

	apir.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminAuth(writer, request)
	})
}

func (thiz *ApiObject) ApiInitAdminca(apir *mux.Router) {
	//test:
	apir.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		thiz.Ermes(writer, "Well, token seems to work!")
	})

	// lessons
	apir.HandleFunc("/lessons/list", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminGetLessons(writer, request)
	})

	apir.HandleFunc("/lessons/save", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminSaveLesson(writer, request)
	})

	apir.HandleFunc("/lessons/delete", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminDeleteLessons(writer, request)
	})

	apir.HandleFunc("/lessons/add", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminAddLessons(writer, request)
	})

	// klasses
	apir.HandleFunc("/klassen/list", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminGetKlasses(writer, request)
	})

	apir.HandleFunc("/klassen/save", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminSaveKlassen(writer, request)
	})

	apir.HandleFunc("/klassen/delete", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminDeleteKlassen(writer, request)
	})

	apir.HandleFunc("/klassen/add", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminAddKlassen(writer, request)
	})

	// timetables
	apir.HandleFunc("/timetables/list", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminGetTimeTables(writer, request)
	})

	apir.HandleFunc("/timetables/add", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminAddTimeTables(writer, request)
	})

	apir.HandleFunc("/timetables/commit", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminCommitTimeTables(writer, request)
	})

	apir.HandleFunc("/gena/list", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminListGena(writer, request)
	})

	apir.HandleFunc("/gena/add", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminAddGena(writer, request)
	})

	apir.HandleFunc("/gena/save", func(writer http.ResponseWriter, request *http.Request) {
		thiz.handlerAdminSaveGena(writer, request)
	})

}
