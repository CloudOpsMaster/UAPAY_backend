package ecom

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	sessionURL string = "https://api.demo.uapay.ua/api/sessions/create"
)

var (
	sessionId string
	jwtKey    = []byte("FJIx7AKc798sQFj8VGALBg==")
)

func getUnixTime() int64 {
	currentTime := time.Now().Unix()
	var uTime int64
	uTime = currentTime + 20
	return uTime
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	iat := getUnixTime()

	// Token jwt Standard Claim Object

	type params struct {
		ClientId string `json:"clientId"`
	}

	type Token struct {
		params
		Iat   int64  `json:"iat"`
		token string `json:"token"`
		jwt.StandardClaims
	}

	var tokenClaim = Token{
		params: params{
			ClientId: "6412",
		},
		Iat:            iat,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			// ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	// Create a new claim with HS256 algorithm and token claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err)
	}

	message := map[string]interface{}{
		"params": map[string]string{
			"clientId": "6412",
		},
		"iat":   iat,
		"token": tokenString,
	}

	// json.NewEncoder(w).Encode(tokenString)
	// json.NewEncoder(w).Encode(message)

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

	var dataToken = result["data"].(map[string]interface{})

	log.Println(result)
	// json.NewEncoder(w).Encode(result)
	json.NewEncoder(w).Encode(dataToken["token"])

	tokenString = dataToken["token"].(string)

	type customClaims struct {
		Id  string `json:"id"`
		Iat string `json:"iat"`
		jwt.StandardClaims
	}

	decodeToken, err := jwt.ParseWithClaims(
		tokenString,
		&customClaims{},
		func(decodeToken *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	claims := decodeToken.Claims.(*customClaims)
	json.NewEncoder(w).Encode(claims.Id)

	sessionId = claims.Id
}

func DemoCreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	iat := getUnixTime()

	message := map[string]interface{}{
		"params": map[string]string{
			"clientId": "6412",
		},
		"iat":   iat,
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
