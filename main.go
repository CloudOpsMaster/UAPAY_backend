package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

const (
	sessionURL       string = "https://api.demo.uapay.ua/api/sessions/create"
	createInvoiceUrl string = "https://api.demo.uapay.ua/api/invoicer/invoices/create"
)

var (
	sessionId string
)

func main() {

	runtime.GOMAXPROCS(2)
	r := mux.NewRouter()

	r.HandleFunc("/create/session", createSession).Methods("GET")
	r.HandleFunc("/create/invoce", createInvoce).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func createSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := map[string]interface{}{
		"params": map[string]string{
			"clientId": "6412",
		},
		"iat":   1611741955,
		"token": "eyJwYXJhbXMiOnsiY2xpZW50SWQiOiI2NDEyIn0sImlhdCI6MTYxMTc0MTk1NSwiYWxnIjoiSFMyNTYifQ.e30.iddIJYnbLyq2pvNAdQGUxI1e4IQ_xu7U169gWiRv4EA",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(sessionURL, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	// map resp.Body
	var result map[string]interface{}

	// Decode result
	json.NewDecoder(resp.Body).Decode(&result)

	// map inside resp.Body Data object
	var data = result["data"].(map[string]interface{})

	log.Println(result["status"])
	log.Println(data["id"])

	json.NewEncoder(w).Encode(data["id"])

	sessionId = data["id"].(string)

}

/*
{
    "amount": 1000,
    "description": "Some book"
}
*/

func createInvoce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var invoiceData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&invoiceData)

	json.NewEncoder(w).Encode(invoiceData)
	json.NewEncoder(w).Encode(invoiceData["description"])

	// geting values amount, description from posts
	var description = invoiceData["description"]
	var amount = invoiceData["amount"]

	//currentTime := time.Now().Unix()
	currentTime := time.Now()
	var externalId string = currentTime.String()

	invoiceMessage := map[string]interface{}{
		"params": map[string]string{
			"sessionId":  sessionId,
			"systemType": "ECOM",
		},
		"data": map[string]interface{}{
			"externalId":  externalId,
			"description": description,
			"amount":      amount,
			"redirectUrl": "https://uapay.ua",
			"type":        "PAY",
			"extraInfo":   "{\"phoneFrom\":\"380971112233\",\"phoneTo\":\"380631112233\",\"cardToId\":\"216f8390-9abc-428d-89d6-7be50183afb5\"}",
		},
		"iat":   1611741955,
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXJhbXMiOnsic2Vzc2lvbklkIjoiZTQ2Zjk5YWQtNWZhNi00Njg2LWE0ZTMtYjdiODhhZjZjM2VhIiwic3lzdGVtVHlwZSI6IkVDT00ifSwiZGF0YSI6eyJleHRlcm5hbElkIjoiMTUwMDM4MzA3NTAwIiwicmV1c2FiaWxpdHkiOmZhbHNlLCJkZXNjcmlwdGlvbiI6ItGC0LXRgdGC0L7QstGL0Lkg0L_Qu9Cw0YJp0LYiLCJhbW91bnQiOjEwMCwicmVkaXJlY3RVcmwiOiJodHRwczovL3VhcGF5LnVhIiwidHlwZSI6IlBBWSIsImV4dHJhSW5mbyI6IntcInBob25lRnJvbVwiOlwiMzgwOTcxMTEyMjMzXCIsXCJwaG9uZVRvXCI6XCIzODA2MzExMTIyMzNcIixcImNhcmRUb0lkXCI6XCIyMTZmODM5MC05YWJjLTQyOGQtODlkNi03YmU1MDE4M2FmYjVcIn0ifX0.5M4zgtmEqfMViuCBigILlzKRGSY6VrmKw-g9CtY7KP8",
	}

	bytesRepresentationInvoice, err := json.Marshal(invoiceMessage)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(createInvoiceUrl, "application/json", bytes.NewBuffer(bytesRepresentationInvoice))
	if err != nil {
		log.Fatalln(err)
	}

	var resultInvoice map[string]interface{}

	// Decode result
	json.NewDecoder(resp.Body).Decode(&resultInvoice)

	// map resultInvoice data
	var paymentPageUrl = resultInvoice["data"].(map[string]interface{})

	//log.Println(resultInvoice)
	log.Println(paymentPageUrl["paymentPageUrl"])

	//json.NewEncoder(w).Encode(resultInvoice)
	json.NewEncoder(w).Encode(paymentPageUrl["paymentPageUrl"])

}
