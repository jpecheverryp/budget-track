package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
    ts, err := template.ParseFiles("./ui/html/pages/home.html")
    if err!=nil{
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = ts.Execute(w, nil)
    if err!=nil{
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func getTransactionView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
    fmt.Fprintf(w, "Showing a single transaction info with id: %d", id)
}

func getTransactionCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show page to add transaction"))
}
func postTransactionCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new transaction"))
}
