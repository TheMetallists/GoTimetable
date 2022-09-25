package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
)

type retLessonList struct {
	Error bool        `json:"error"`
	Items []retLesson `json:"items"`
}

type retLesson struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Cabinet string `json:"cabinet"`
}

func (thiz *ApiObject) handlerAdminGetLessons(w http.ResponseWriter, r *http.Request) {
	ret := retLessonList{
		Error: false,
		Items: []retLesson{},
	}

	lines := []database.DB_Lessons{}

	thiz.DH.DB.Find(&lines)

	for _, item := range lines {
		ret.Items = append(ret.Items, retLesson{
			ID:      item.ID,
			Name:    item.Name,
			Cabinet: item.Cabinet,
		})
	}

	jsn, err := json.MarshalIndent(ret, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
