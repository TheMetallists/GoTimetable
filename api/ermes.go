package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponce struct {
	Error bool   `json:"error"`
	ErMes string `json:"error_msg"`
}

func (thiz *ApiObject) Ermes(w http.ResponseWriter, ermes string) {
	erre := errorResponce{
		Error: true,
		ErMes: ermes,
	}
	jsn, err := json.MarshalIndent(erre, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsn)
}
