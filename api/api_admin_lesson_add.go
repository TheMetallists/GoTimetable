package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
)

func (thiz *ApiObject) handlerAdminAddLessons(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	if len(r.MultipartForm.Value["name"]) != 1 {
		thiz.Ermes(w, "No lesson name provided!")
		return
	}

	lesEntry := database.DB_Lessons{
		Name:    r.MultipartForm.Value["name"][0],
		Cabinet: "1111A",
	}

	thiz.DH.DB.Create(&lesEntry)

	jsn, err := json.MarshalIndent(map[string]interface{}{
		"error": false,
		"id":    lesEntry.ID,
	}, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
