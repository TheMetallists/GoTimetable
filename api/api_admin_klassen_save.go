package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminSaveKlassen(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	if len(r.MultipartForm.Value["id"]) != 1 {
		thiz.Ermes(w, "No id supplied!s")
		return
	}

	id, err := strconv.ParseInt(r.MultipartForm.Value["id"][0], 10, 64)

	if err != nil {
		thiz.Ermes(w, "Invalid id!")
		return
	}

	lines := database.DB_Klass{}

	ctr := int64(0)
	thiz.DH.DB.First(&lines, id).Count(&ctr)

	if ctr == 0 {
		thiz.Ermes(w, "Nothing found!")
		return
	}

	if len(r.MultipartForm.Value["letter"]) > 0 {
		lines.Letter = r.MultipartForm.Value["letter"][0]
	}
	if len(r.MultipartForm.Value["number"]) > 0 {

		numb, err := strconv.ParseInt(r.MultipartForm.Value["number"][0], 10, 8)

		if err != nil || numb < 1 || numb > 11 {
			thiz.Ermes(w, "Invalid klass number!")
			return
		}
		lines.Number = uint8(numb)
	}

	thiz.DH.DB.Save(lines)

	jsn, err := json.MarshalIndent(map[string]bool{
		"error": false,
	}, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
	return
}
