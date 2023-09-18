package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const clientID = ""
const clientSecret = ""

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	httpClient := http.Client{}

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error in parsing query: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		code := r.FormValue("code")

		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			log.Println("Error in HTTP request creation: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req.Header.Set("accept", "application/json")
		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("Error in sending HTTP request: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// Parse the request body into the `OAuthAccessResponse` struct
		var t OAuthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			log.Println("Error in parsing JSON response: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
