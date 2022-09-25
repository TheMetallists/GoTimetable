package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminSaveGena(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	inputId := uint(0)

	if len(r.MultipartForm.Value["id"]) != 1 {
		thiz.Ermes(w, "No parallel provided!")
		return
	} else {
		eid, err := strconv.ParseInt(r.MultipartForm.Value["id"][0], 10, 32)
		if err == nil {
			inputId = uint(eid)
		} else {
			thiz.Ermes(w, "Invalid parallel!")
			return
		}
	}

	inputHours := uint8(0)

	if len(r.MultipartForm.Value["hours"]) != 1 {
		thiz.Ermes(w, "No lesson provided!")
		return
	} else {
		numb, err := strconv.ParseInt(r.MultipartForm.Value["hours"][0], 10, 32)
		if err == nil {
			inputHours = uint8(numb)
		} else {
			thiz.Ermes(w, "Invalid lesson!")
			return
		}
	}

	if inputId < 1 || inputHours < 0 {
		thiz.Ermes(w, "ID or hours wrong.")
		return
	}

	lines := database.DB_LessonGeneratorEntry{}

	ctr := int64(0)
	thiz.DH.DB.First(&lines, inputId).Count(&ctr)

	if ctr == 0 {
		thiz.Ermes(w, "Nothing found!")
		return
	}

	lines.HoursPerWeek = inputHours

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
