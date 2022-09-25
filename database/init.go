package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

type initLesson struct {
	Name    string
	Cabinet string
	DBID    uint
}

func insertTestingData(_dh *DatabaseHolder) {
	_dh.DB.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Initializing DB...")

		lessons := [...]initLesson{
			{Name: "Матан", Cabinet: "4"},
			{Name: "Шизика", Cabinet: "42"},
			{Name: "Химия", Cabinet: "16"},
			{Name: "Биология", Cabinet: "8"},
			{Name: "ОБЖорка", Cabinet: "52"},
			{Name: "Музло", Cabinet: "27"},
			{Name: "Русичи", Cabinet: "49"},
			{Name: "Литерча", Cabinet: "37"},
			{Name: "Физуха", Cabinet: "38"},
			{Name: "Труд", Cabinet: "-1"},
			{Name: "English", Cabinet: "44"},
			{Name: "Инф-ра", Cabinet: "1111A"},
		}

		for i, itm := range lessons {
			log.Println("[DB:INIT] Creating lesson: ", itm.Name, " in room: ", itm.Cabinet)
			lsn := DB_Lessons{
				Name:    itm.Name,
				Cabinet: itm.Cabinet,
			}
			ret := tx.Create(&lsn)
			if ret.Error != nil {
				log.Println("Error creating lessons: ", ret.Error)
			}
			lessons[i].DBID = lsn.ID
		}

		for kn := 1; kn <= 11; kn++ {
			klt := [...]string{"А", "Б", "В", "Г"}
			for _, kl := range klt {
				log.Println("[DB:INIT] Creating class: ", kn, kl)

				klass := DB_Klass{
					Number: uint8(kn),
					Letter: kl,
				}
				retk := tx.Create(&klass)
				if retk.Error != nil {
					log.Fatalln("Error creating DEMO KLASS: ", retk.Error)
				}

				for wd := 1; wd <= 6; wd++ {
					numLessont := rand.Intn(4) + 4
					for nl := 0; nl < numLessont; nl++ {
						lesEntry := DB_LessonEntry{
							Klass:   klass.ID,
							Weekday: uint8(wd),
							Lesson:  lessons[rand.Intn(len(lessons))].DBID,
						}

						tx.Create(&lesEntry)
					}
				}

			}
		}

		return nil
	})

}

func DatabaseInit() *DatabaseHolder {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open(dir+"/databank.db3"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot open database: ", err)
	}

	err = db.AutoMigrate(&DB_Klass{})
	if err != nil {
		log.Fatalln("Migration error: ", err)
	}
	err = db.AutoMigrate(&DB_Lessons{})
	if err != nil {
		log.Fatalln("Migration error: ", err)
	}
	err = db.AutoMigrate(&DB_LessonEntry{})
	if err != nil {
		log.Fatalln("Migration error: ", err)
	}
	err = db.AutoMigrate(&DB_Options{})
	if err != nil {
		log.Fatalln("Migration error: ", err)
	}

	err = db.AutoMigrate(&DB_LessonGeneratorEntry{})
	if err != nil {
		log.Fatalln("Migration error: ", err)
	}

	dh := &DatabaseHolder{
		db,
		nil,
	}

	if dh.GetOptionString("db.inited", "no") == "no" {
		insertTestingData(dh)
		dh.SetOptionString("db.inited", "yes")
	}
	dh.GetJwtTokenKey()

	return dh
}
