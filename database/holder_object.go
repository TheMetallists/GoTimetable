package database

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
	"log"
)

type DatabaseHolder struct {
	DB       *gorm.DB
	tokenKey *string
}

func (thiz *DatabaseHolder) GetOptionString(_name, _default string) string {
	fnd := DB_Options{
		Model: gorm.Model{},
		Name:  "",
		Value: _default,
	}

	ret := thiz.DB.Where(&DB_Options{
		Name: _name,
	}).Find(&fnd)

	if ret.Error != nil {
		return _default
	}

	return fnd.Value
}

func (thiz *DatabaseHolder) SetOptionString(_name, _value string) error {
	rc := thiz.DB.
		Model(&DB_Options{}).
		Where(&DB_Options{
			Name: _name,
		}).Update("value", _value)

	if rc.RowsAffected < 1 {
		thiz.DB.Create(&DB_Options{
			Name:  _name,
			Value: _value,
		})
	}

	return rc.Error
}

func (thiz *DatabaseHolder) GetJwtTokenKey() string {
	if thiz.tokenKey != nil {
		return *thiz.tokenKey
	}

	key := thiz.GetOptionString("auth.tokenkey", "null")

	if key == "null" {
		kex := make([]byte, 256)

		_, err := rand.Read(kex)
		if err != nil {
			log.Fatalln("Error generating crypto key!", err)
		}

		key = base64.StdEncoding.EncodeToString(kex)
		thiz.SetOptionString("auth.tokenkey", key)
	}

	thiz.tokenKey = &key

	thiz.PopulateAdminPassword()
	return key
}
func (thiz *DatabaseHolder) PopulateAdminPassword() {
	kex := make([]byte, 9)

	_, err := rand.Read(kex)
	if err != nil {
		log.Fatalln("Error generating crypto key!", err)
	}

	key := base64.StdEncoding.EncodeToString(kex)
	log.Println("NEW ADMIN PASSWORD: ", key)

	adminPassSalted := pbkdf2.Key([]byte(key), []byte(*thiz.tokenKey), 12000, 64, sha256.New)
	adminPassSaltedb64 := base64.StdEncoding.EncodeToString(adminPassSalted)

	thiz.SetOptionString("auth.adminpass", adminPassSaltedb64)
}

func (thiz *DatabaseHolder) GetAdminPassword() []byte {
	encpw := thiz.GetOptionString("auth.adminpass", "null")

	if encpw == "null" {
		log.Fatalln("Admin password not initialized!")
	}

	ret, err := base64.StdEncoding.DecodeString(encpw)
	if err != nil {
		log.Fatalln("Cannot decode admin's password key", err)
	}

	return ret
}
