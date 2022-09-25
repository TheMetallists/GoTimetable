package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
)

type retKlassenList struct {
	Error bool       `json:"error"`
	Items []retKlass `json:"items"`
}

type retKlass struct {
	ID           uint   `json:"id"`
	Number       uint8  `json:"number"`
	Letter       string `json:"letter"`
	CountLessons int64  `json:"countLessons"`
}

func (thiz *ApiObject) handlerAdminGetKlasses(w http.ResponseWriter, r *http.Request) {
	ret := retKlassenList{
		Error: false,
		Items: []retKlass{},
	}

	lines := []database.DB_Klass{}

	thiz.DH.DB.Order("number ASC, letter").Find(&lines)

	for _, item := range lines {
		cntLes := int64(0)
		thiz.DH.DB.
			Model(database.DB_LessonEntry{}).
			Where("klass = ?", item.ID).
			Count(&cntLes)

		ret.Items = append(ret.Items, retKlass{
			ID:           item.ID,
			Number:       item.Number,
			Letter:       item.Letter,
			CountLessons: cntLes,
		})
	}

	jsn, err := json.MarshalIndent(ret, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
