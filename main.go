package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/Maks0123/UAPAY_backend/ecom"
	"github.com/gorilla/mux"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {

	runtime.GOMAXPROCS(1)
	HandleFunction()

}

func HandleFunction() {

	// var port = "5000"

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/demo/create/session", ecom.DemoCreateSession).Methods("GET")
	r.HandleFunc("/create/session", ecom.CreateSession).Methods("GET")
	r.HandleFunc("/demo/create/invoce", ecom.DemoCreateInvoce).Methods("POST")
	r.HandleFunc("/create/invoce", ecom.CreateInvoce).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))

}

func Index(w http.ResponseWriter, r *http.Request) {

	data := ViewData{
		Title:   "World Cup",
		Message: "FIFA will never regret it",
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}
