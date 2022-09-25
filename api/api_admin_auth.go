package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"net/http"
	"time"
)

type jwtClaims struct {
	User string `json:"y"`
	jwt.StandardClaims
}

type tokenResponce struct {
	Error bool   `json:"error"`
	Token string `json:"token"`
}

func (thiz *ApiObject) handlerAdminAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Println("RQST ERROR: ", err)
		thiz.Ermes(w, "Error: no password supplied!")
		return
	}

	passw := r.MultipartForm.Value["password"]
	if len(passw) != 1 {
		thiz.Ermes(w, "Error: no password supplied!")
		return
	}

	adminPassSalted := pbkdf2.Key([]byte(passw[0]), []byte(thiz.DH.GetJwtTokenKey()), 12000, 64, sha256.New)

	if bytes.Compare(adminPassSalted, thiz.DH.GetAdminPassword()) == 0 {
		// password ok

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
			"admin",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
				Issuer:    "primaryServer",
			},
		})
		tkn, err := jwtToken.SignedString([]byte(thiz.DH.GetJwtTokenKey()))
		if err != nil {
			log.Println("Error generating token: ", err)
			thiz.Ermes(w, "Error: Internal server error!")
			return
		}

		resp := tokenResponce{
			Error: false,
			Token: tkn,
		}

		jsn, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(jsn)
		return
	}

	thiz.Ermes(w, "Invalid password!")

}
