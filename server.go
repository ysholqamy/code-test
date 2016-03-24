package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ysholqamy/code-test/event"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// should be moved to event package as singelton?
var globalEventManager = event.NewManager()

func generateUniqueSessionId() string {
	rand.Seed(time.Now().UnixNano())
	const count = 5
	const intByteSize = 8 // will be using int64
	b := make([]byte, count*intByteSize)
	for i := 0; i < count; i++ {
		// add random int64 to buffer
		binary.PutVarint(b, rand.Int63())
	}
	// return base64 encoding of the generated buffer
	return base64.StdEncoding.EncodeToString(b)
}

func prepareRequest(w http.ResponseWriter, r *http.Request) error {
	log.Println("preparing request")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("Only accept POST, got %s", r.Method)
	}

	//Set headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return nil
}

func handleSessionRequest(w http.ResponseWriter, r *http.Request) {
	err := prepareRequest(w, r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// will be rendering json, set header.
	w.Header().Set("Content-Type", "application/json; charset=UTF8")

	resBody := map[string]string{"sessionId": generateUniqueSessionId()}
	js, err := json.Marshal(resBody)

	//duplication -.-
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(js)

}

func handleEventRequest(w http.ResponseWriter, r *http.Request) {
	err := prepareRequest(w, r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Prepare JSON decoder
	decoder := json.NewDecoder(r.Body)

	ch := make(chan event.Result)

	log.Println("processing event...")
	go globalEventManager.RegisterEvent(decoder, ch)
	log.Println("done processing...")
	result := <-ch
	log.Println("got result out of channel")
	if result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest) // or InternalServerError maybe?
		return
	}

	log.Println(result.Message)
}

func main() {
	http.HandleFunc("/", handleEventRequest)
	http.HandleFunc("/session", handleSessionRequest)
	log.Println("listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
