package api

import (
	"dislexiad/database"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func getNumber(r *http.Request, name string) (int64, error) {
	if len(r.MultipartForm.Value[name]) != 1 {
		return 0, errors.New(name + " is not provided")
	}

	numb, err := strconv.ParseInt(r.MultipartForm.Value[name][0], 10, 8)

	if err != nil || numb < 1 {
		return 0, errors.New(name + ": wrong number specified")
	}

	return numb, nil
}

func (thiz *ApiObject) handlerAdminAddTimeTables(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		thiz.Ermes(w, "Form error!")
		return
	}

	klass, err := getNumber(r, "klass")
	if err != nil {
		thiz.Ermes(w, err.Error())
		return
	}

	lesson, err := getNumber(r, "lesson")
	if err != nil {
		thiz.Ermes(w, err.Error())
		return
	}

	weekday, err := getNumber(r, "weekday")
	if err != nil {
		thiz.Ermes(w, err.Error())
		return
	}

	//TODO: verify data further!
	lesEntry := database.DB_LessonEntry{
		Klass:   uint(klass),
		Weekday: uint8(weekday),
		Lesson:  uint(lesson),
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
