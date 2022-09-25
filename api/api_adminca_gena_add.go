package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (thiz *ApiObject) handlerAdminAddGena(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	inputParallel := uint8(0)

	if len(r.MultipartForm.Value["parallel"]) != 1 {
		thiz.Ermes(w, "No parallel provided!")
		return
	} else {
		numb, err := strconv.ParseInt(r.MultipartForm.Value["parallel"][0], 10, 32)
		if err == nil {
			inputParallel = uint8(numb)
		} else {
			thiz.Ermes(w, "Invalid parallel!")
			return
		}
	}

	inputLesson := uint(0)

	if len(r.MultipartForm.Value["lesson"]) != 1 {
		thiz.Ermes(w, "No lesson provided!")
		return
	} else {
		numb, err := strconv.ParseInt(r.MultipartForm.Value["lesson"][0], 10, 32)
		if err == nil {
			inputLesson = uint(numb)
		} else {
			thiz.Ermes(w, "Invalid lesson!")
			return
		}
	}

	if inputLesson < 1 || inputParallel < 1 {
		thiz.Ermes(w, "lesson or parallel wrong.")
		return
	}

	lesson := database.DB_Lessons{}
	cnt := int64(0)
	thiz.DH.DB.
		Model(database.DB_Lessons{}).
		Where("id = ?", inputLesson).
		First(&lesson).Count(&cnt)

	if cnt < 1 {
		thiz.Ermes(w, "Lesson not found!")
		return
	}

	klasses := make([]database.DB_Klass, 0)
	thiz.DH.DB.
		Model(database.DB_Klass{}).
		Where("number = ?", inputParallel).
		Find(&klasses).Count(&cnt)

	for _, kix := range klasses {
		thiz.DH.DB.
			Model(database.DB_LessonGeneratorEntry{}).
			Where("klass = ? AND lesson = ?", kix.ID, lesson.ID).
			Count(&cnt)
		if cnt < 1 {
			line := database.DB_LessonGeneratorEntry{
				Klass:        kix.ID,
				Lesson:       lesson.ID,
				HoursPerWeek: 1,
			}
			thiz.DH.DB.Create(&line)
		}
	}

	jsn, err := json.MarshalIndent(map[string]interface{}{
		"error": false,
	}, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)

}
