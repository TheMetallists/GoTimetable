package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminSaveLesson(w http.ResponseWriter, r *http.Request) {
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

	lines := database.DB_Lessons{}

	ctr := int64(0)
	thiz.DH.DB.First(&lines, id).Count(&ctr)

	if ctr == 0 {
		thiz.Ermes(w, "Nothing found!")
		return
	}

	if len(r.MultipartForm.Value["name"]) > 0 {
		lines.Name = r.MultipartForm.Value["name"][0]
	}
	if len(r.MultipartForm.Value["cabinet"]) > 0 {
		lines.Cabinet = r.MultipartForm.Value["cabinet"][0]
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
