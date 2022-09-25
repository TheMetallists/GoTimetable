package database

import "gorm.io/gorm"

type DB_Klass struct {
	gorm.Model
	Number uint8
	Letter string
}

type DB_Lessons struct {
	gorm.Model
	Name    string
	Cabinet string
}

type DB_LessonEntry struct {
	gorm.Model
	Klass   uint
	Weekday uint8
	Lesson  uint
}

type DB_LessonGeneratorEntry struct {
	gorm.Model
	Klass        uint
	Lesson       uint
	HoursPerWeek uint8
}

type DB_Options struct {
	gorm.Model
	Name  string
	Value string
}
