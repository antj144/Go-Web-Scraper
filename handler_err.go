package main


import "net/http" 


func handlerErr(w, http.ResponseWriter, r *http.Request){
	respondWithError00(w, 400, "Somthing went wrong")
}
