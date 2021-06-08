package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
)

func main() {
	debug.SetPanicOnFault(true)
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", country)

	fmt.Printf("Starting server at port 80\n")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["name"]

	name := "guest"

	if ok {

		name = keys[0]
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
}

func country(w http.ResponseWriter, r *http.Request) {
	debug.SetPanicOnFault(true)
	url := "http://country.io/capital.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	fmt.Fprintf(w, responseString)
}
