package main

import "net/http"

type myStruct struct{
	Name string;
}

func Handlerfunc(w http.ResponseWriter, r *http.Request) {
	s := myStruct{"Fucj=k"}
	RespondwithJson(w,200, s)
}

func Handlererr(w http.ResponseWriter, r *http.Request) {
	Respondwitherror(w, 400, "Kuch toh Gadbad H")
}