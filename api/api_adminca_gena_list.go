package api

import (
	"dislexiad/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type retGenaLesson struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Cabinet string `json:"cabinet"`
	WasUsed bool   `json:"used"`
}

type retGenaList struct {
	Error       bool                   `json:"error"`
	Parallels   []int                  `json:"parallels"`
	Lessons     []retGenaLesson        `json:"lessons"`
	CurParallel retGenaCurrentParallel `json:"currentParallel"`
}

type retGenaEntry struct {
	ID           uint  `json:"id"`
	Klass        uint  `json:"klass"`
	Lesson       uint  `json:"lesson"`
	HoursPerWeek uint8 `json:"hoursPerWeek"`
}

type retGenaCurrentParallel struct {
	Number   uint8                          `json:"number"`
	Klasses  []retKlass                     `json:"klasses"`
	Elements map[uint]map[uint]retGenaEntry `json:"elements"`
}

func (thiz *ApiObject) handlerAdminListGena(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println(err)
		thiz.Ermes(w, "Wrong form error!")
		return
	}

	inputParallel := uint8(0)

	if len(r.MultipartForm.Value["parallel"]) != 1 {
		inputParallel = 0
	} else {
		numb, err := strconv.ParseInt(r.MultipartForm.Value["parallel"][0], 10, 32)
		if err == nil {
			inputParallel = uint8(numb)
		}
	}

	retn := retGenaList{
		Error:     false,
		Parallels: make([]int, 0),
		Lessons:   make([]retGenaLesson, 0),
		CurParallel: retGenaCurrentParallel{
			Number:   0,
			Klasses:  make([]retKlass, 0),
			Elements: make(map[uint]map[uint]retGenaEntry),
		},
	}

	// find parallels
	lines := []database.DB_Klass{}
	thiz.DH.DB.
		Model(database.DB_Klass{}).
		Group("number").Find(&lines)
	for _, kls := range lines {
		retn.Parallels = append(retn.Parallels, int(kls.Number))
	}

	// loading lessons
	lessns := []database.DB_Lessons{}
	thiz.DH.DB.Find(&lessns)
	for _, lsn := range lessns {
		retn.Lessons = append(retn.Lessons, retGenaLesson{
			ID:      lsn.ID,
			Name:    lsn.Name,
			Cabinet: lsn.Cabinet,
			WasUsed: false,
		})
	}

	// currentParallel
	if len(retn.Parallels) < 1 {
		jsn, err := json.MarshalIndent(retn, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(jsn)
		return
	} else {
		bHasFound := false
		for _, itm := range retn.Parallels {
			if itm == int(inputParallel) {
				bHasFound = true
			}
		}

		if !bHasFound {
			inputParallel = uint8(retn.Parallels[0])
		}
	}

	retn.CurParallel.Number = inputParallel

	lines2 := []database.DB_Klass{}
	thiz.DH.DB.
		Model(database.DB_Klass{}).
		Where("number = ?", inputParallel).
		Find(&lines2)

	klassIDs := make([]uint, 0)

	for _, kls := range lines2 {
		klassIDs = append(klassIDs, kls.ID)
		retn.CurParallel.Klasses = append(retn.CurParallel.Klasses, retKlass{
			ID:           kls.ID,
			Number:       kls.Number,
			Letter:       kls.Letter,
			CountLessons: 0,
		})
	}

	genas := []database.DB_LessonGeneratorEntry{}
	thiz.DH.DB.
		Model(database.DB_LessonGeneratorEntry{}).
		Where("klass IN (?)", klassIDs).
		Find(&genas)

	usedLessons := make(map[uint]bool)

	for _, gena := range genas {
		usedLessons[gena.Lesson] = true
		if retn.CurParallel.Elements[gena.Klass] == nil {
			retn.CurParallel.Elements[gena.Klass] = make(map[uint]retGenaEntry)
		}
		retn.CurParallel.Elements[gena.Klass][gena.Lesson] = retGenaEntry{
			ID:           gena.ID,
			Klass:        gena.Klass,
			Lesson:       gena.Lesson,
			HoursPerWeek: gena.HoursPerWeek,
		}
	}

	for idl, _ := range retn.Lessons {
		if _, ok := usedLessons[retn.Lessons[idl].ID]; ok {
			retn.Lessons[idl].WasUsed = true
		}
	}

	// inputParallel

	jsn, err := json.MarshalIndent(retn, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
