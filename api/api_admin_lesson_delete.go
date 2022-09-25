package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminDeleteLessons(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	if len(r.MultipartForm.Value["id"]) < 1 {
		thiz.Ermes(w, "No ids supplied!")
		return
	}

	affectedL := 0
	affectedLE := 0
	for _, strid := range r.MultipartForm.Value["id"] {
		id, err := strconv.ParseInt(strid, 10, 64)

		if err != nil {
			thiz.Ermes(w, "Invalid id!")
			return
		}

		line := database.DB_Lessons{}

		ctr := int64(0)
		thiz.DH.DB.First(&line, id).Count(&ctr)

		if ctr == 0 {
			continue
		}

		lents := []database.DB_LessonEntry{}
		// remove all lessons in this classroom
		thiz.DH.DB.
			Model(database.DB_LessonEntry{}).
			Where("lesson = ?", line.ID).
			Find(&lents)

		for _, lent := range lents {
			log.Println("Deleting lesson entry:", lent)
			//TODO add a crawler to delete soft-deleted items after a certain period of time.
			thiz.DH.DB.Delete(&lent)
			_ = lent
			affectedLE += 1
		}

		log.Println("Deleting lesson: ", line)
		thiz.DH.DB.Delete(&line)
		affectedL += 1
	}

	jsn, err := json.MarshalIndent(map[string]interface{}{
		"error":      false,
		"affectedL":  affectedL,
		"affectedLE": affectedLE,
	}, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
	return
}
