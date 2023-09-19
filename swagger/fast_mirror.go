package main

import (
	"encoding/json"
	"fmt"
	"github.com/hemantkgupta/my-go-work/goplay/mirrors"
	"log"
	"net/http"
	"time"
)

// $docker run -p 80:8080 -e SWAGGER_JSON=/app/open.json -v <code-dir>/my-go-work/swagger:/app  swaggerapi/swagger-ui
type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter,
		r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Starting server on port", port)
	log.Fatal(server.ListenAndServe())
}
func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}
