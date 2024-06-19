package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondwithJson(w http.ResponseWriter, code int, payload interface{}) {
	val,err := json.Marshal(payload)
	if err!=nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type","application/json")
	w.WriteHeader(code)
	w.Write(val)
}

func Respondwitherror(w http.ResponseWriter, code int, msg string){
	if code>499{
		log.Println("Responding with 5XX error:",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}

	RespondwithJson(w,code,errResponse{msg})
}