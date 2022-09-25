package api

import (
	"dislexiad/database"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type retTimeTablesList struct {
	Error   bool                     `json:"error"`
	Items   map[uint8][]retTimeTable `json:"items"`
	Klass   retKlass                 `json:"klass"`
	Lessons map[uint]retLesson       `json:"lessons"`
}

type retTimeTable struct {
	ID      uint  `json:"id"`
	Weekday uint8 `json:"weekday"` // we now the klass id.
	Lesson  uint  `json:"lesson"`
}

func (thiz *ApiObject) handlerAdminGetTimeTables(w http.ResponseWriter, r *http.Request) {
	ret := retTimeTablesList{
		Error:   false,
		Items:   make(map[uint8][]retTimeTable),
		Klass:   retKlass{},
		Lessons: make(map[uint]retLesson),
	}

	err := r.ParseMultipartForm(1024)
	if err != nil {
		thiz.Ermes(w, "No ids supplied!")
		return
	}

	if len(r.MultipartForm.Value["id"]) < 1 {
		thiz.Ermes(w, "No ids supplied!")
		return
	}
	id, err := strconv.ParseInt(r.MultipartForm.Value["id"][0], 10, 64)

	if err != nil {
		thiz.Ermes(w, "Invalid id!")
		return
	}

	/// loading klass
	klass := &database.DB_Klass{
		Model:  gorm.Model{},
		Number: 0,
		Letter: "",
	}

	kctr := int64(0)
	thiz.DH.DB.First(&klass, id).Count(&kctr)

	if kctr < 1 {
		thiz.Ermes(w, "Klass not found!")
		return
	}

	ret.Klass = retKlass{
		ID:           klass.ID,
		Number:       klass.Number,
		Letter:       klass.Letter,
		CountLessons: 0,
	}

	/// loading lesson list

	lesns := []database.DB_Lessons{}
	thiz.DH.DB.Find(&lesns)

	for _, litem := range lesns {
		ret.Lessons[litem.ID] = retLesson{
			ID:      litem.ID,
			Name:    litem.Name,
			Cabinet: litem.Cabinet,
		}
	}

	// load klass-lesson relations

	lents := []database.DB_LessonEntry{}
	thiz.DH.DB.
		Model(database.DB_LessonEntry{}).
		Where("klass = ?", ret.Klass.ID).
		Order("weekday ASC, ID").
		Find(&lents)

	for _, lent := range lents {
		ret.Items[lent.Weekday] = append(ret.Items[lent.Weekday], retTimeTable{
			ID:      lent.ID,
			Weekday: lent.Weekday,
			Lesson:  lent.Lesson,
		})
	}

	jsn, err := json.MarshalIndent(ret, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
