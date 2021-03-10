package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/Maks0123/UAPAY_backend/ecom"
	"github.com/gorilla/mux"
)

func main() {

	runtime.GOMAXPROCS(1)
	HandleFunction()

}

func HandleFunction() {

	var port = "5000"

	// port := os.Getenv("PORT")

	/*	if port == "" {
			log.Fatal("$PORT must be set")
		}
	*/

	r := mux.NewRouter()

	r.HandleFunc("/", TestFunc).Methods("GET")
	r.HandleFunc("/demo/create/session", ecom.DemoCreateSession).Methods("GET")
	r.HandleFunc("/create/session", ecom.CreateSession).Methods("GET")
	r.HandleFunc("/demo/create/invoce", ecom.DemoCreateInvoce).Methods("POST")
	r.HandleFunc("/create/invoce", ecom.CreateInvoce).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))

}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	var hello string = "Hello GO"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
