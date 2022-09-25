package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminAddKlassen(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	if len(r.MultipartForm.Value["letter"]) != 1 {
		thiz.Ermes(w, "No klass letter provided!")
		return
	}

	if len(r.MultipartForm.Value["number"]) != 1 {
		thiz.Ermes(w, "No klass number provided!")
		return
	}

	numb, err := strconv.ParseInt(r.MultipartForm.Value["number"][0], 10, 8)

	if err != nil || numb < 1 || numb > 11 {
		thiz.Ermes(w, "Invalid klass number!")
		return
	}

	klasEntry := database.DB_Klass{
		Number: uint8(numb),
		Letter: r.MultipartForm.Value["letter"][0],
	}

	thiz.DH.DB.Create(&klasEntry)

	jsn, err := json.MarshalIndent(map[string]interface{}{
		"error": false,
		"id":    klasEntry.ID,
	}, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
