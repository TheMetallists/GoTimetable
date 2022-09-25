package api

import (
	"crypto/sha1"
	"dislexiad/database"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
{"1А":[{"day":1,"lessons":[{"lesson":"Русичи","cab":49},{"lesson":"ОБЖорка","cab":52}]},{"lesson":"Труд","cab":-1},{"lesson":"Химия","cab":16},{"lesson":"Литерча","cab":37},{"lesson":"Инф-ра","cab":"1111А"}],"notice":""},{"day":2,"lessons":[{"lesson":"ОБЖорка","cab":52},{"lesson":"Химия","cab":16},{"lesson":"Биология","cab":8},{"lesson":"English","cab":44}],"notice":""},{"day":3,"lessons":[{"lesson":"Инф-ра","cab":"1111А"},{"lesson":"Биология","cab":8},{"lesson":"Инф-ра","cab":"1111А"},{"lesson":"Шизика","cab":42},{"lesson":"Русичи","cab":49},{"lesson":"Труд","cab":-1},{"lesson":"Музло","cab":27}],"notice":""},{"day":4,"lessons":[{"lesson":"Шизика","cab":42},{"lesson":"Матан","cab":4},{"lesson":"English","cab":44},{"lesson":"Русичи","cab":49},{"lesson":"Литерча","cab":37},{"lesson":"Инф-ра","cab":"1111А"},

*/

type telekLesson struct {
	Lesson string `json:"lesson"`
	Kab    string `json:"kab"`
}

type teleklassenItem struct {
	Day     uint8         `json:"day"`
	Lessons []telekLesson `json:"lessons"`
}

type telekResponce struct {
	Error   bool                         `json:"error"`
	Klassen map[string][]teleklassenItem `json:"klassen"`
	Hash    string                       `json:"hash"`
}

func (thiz *ApiObject) handlerTelek(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	ret := telekResponce{
		Error:   false,
		Klassen: make(map[string][]teleklassenItem),
	}

	// loading lessons
	lessons := make(map[uint]database.DB_Lessons)
	lsl := []database.DB_Lessons{}
	thiz.DH.DB.Find(&lsl)
	for _, lsi := range lsl {
		lessons[lsi.ID] = lsi
	}

	//search classes
	klassen := []database.DB_Klass{}
	thiz.DH.DB.Order("number ASC, letter ASC").Find(&klassen)

	for _, kls := range klassen {
		tki := []teleklassenItem{}
		lEntries := []database.DB_LessonEntry{}
		thiz.DH.DB.Where("klass = ?", kls.ID).Find(&lEntries)

		for wd := uint8(0); wd < 6; wd++ {
			tki = append(tki, teleklassenItem{
				Day:     wd + 1,
				Lessons: []telekLesson{},
			})
		}

		for _, ken := range lEntries {
			tki[ken.Weekday-1].Lessons = append(tki[ken.Weekday-1].Lessons, telekLesson{
				Lesson: lessons[ken.Lesson].Name,
				Kab:    lessons[ken.Lesson].Cabinet,
			})
		}

		ret.Klassen[fmt.Sprintf("%d%s", kls.Number, kls.Letter)] = tki
	}

	// generating hash for setState
	jsn1, err := json.MarshalIndent(ret.Klassen, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	h := sha1.New()
	h.Write(jsn1)
	ret.Hash = hex.EncodeToString(h.Sum(nil))

	jsn, err := json.MarshalIndent(ret, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
