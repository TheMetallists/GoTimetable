package api

import (
	"dislexiad/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hacttInputTable struct {
	KlassID int                      `json:"kid"`
	Items   map[uint8][]retTimeTable `json:"items"`
}

/*
type retTimeTable struct {
	ID      uint  `json:"id"`
	Weekday uint8 `json:"weekday"` // we now the klass id.
	Lesson  uint  `json:"lesson"`
}
*/

func (thiz *ApiObject) handlerAdminCommitTimeTables(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		thiz.Ermes(w, "Cannot read request body!")
		return
	}

	input := hacttInputTable{}

	err = json.Unmarshal(data, &input)
	if err != nil {
		thiz.Ermes(w, "Invalid json.")
		return
	}

	klass := &database.DB_Klass{}

	kctr := int64(0)
	thiz.DH.DB.First(&klass, input.KlassID).Count(&kctr)

	if kctr < 1 {
		thiz.Ermes(w, "Klass not found!")
		return
	}

	// creating a list of ids of existing items to remove unwanted items

	lents := []database.DB_LessonEntry{}
	thiz.DH.DB.
		Model(database.DB_LessonEntry{}).
		Where("klass = ?", klass.ID).
		Order("weekday ASC, ID").
		Find(&lents)

	entryIds := []uint{}

	for _, kent := range lents {
		entryIds = append(entryIds, kent.ID)
	}
	log.Println("Found KENTS: ", entryIds)

	for _, itemList := range input.Items {
		for _, item := range itemList {
			lentry := database.DB_LessonEntry{}
			kctr := int64(0)
			thiz.DH.DB.First(&lentry, item.ID).Count(&kctr)

			if kctr > 0 {
				lentry.Weekday = item.Weekday
				lentry.Lesson = item.Lesson
				thiz.DH.DB.Save(&lentry)

				// entryIds
				for i, v := range entryIds {
					if v == lentry.ID {
						entryIds = append(entryIds[:i], entryIds[i+1:]...)
					}
				}
			}

		}
	}

	// removing unnecessary items
	for _, lenid := range entryIds {
		lentry := database.DB_LessonEntry{}
		kctr := int64(0)
		fmt.Println("Deleting: ", lenid)
		thiz.DH.DB.First(&lentry, lenid).Count(&kctr)

		if kctr > 0 {
			thiz.DH.DB.Delete(&lentry)
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
