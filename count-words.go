// count words in sent text file to web server
// there should be middleware to measure and log
// execution time
// text file should be cleaned from commas and dottes 
// before counting
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// handleFuncs homePage func executed when url will be http://localhost:4000
func countWords(w http.ResponseWriter, r *http.Request) {
	log.Println("start counting")
	// check if method POST
	if r.Method != "POST" {
		fmt.Fprint(w, "it should be POST")
		return
	}
	// read sent file
	data, err := ioutil.ReadAll(r.Body)
	// do not forget to close it
	defer r.Body.Close()
	// check for err
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	words := make(map[string]int)
	// remove all commas
	text := strings.Replace(string(data), ",", "", -1)
	// remove all dots
	text = strings.Replace(text, ".", "", -1)
	// create slice of all words in text
	allwords := strings.Split(text, " ")
	for _, val := range allwords {
		words[val]++
	}
	// write to our response (data received by Browser)
	fmt.Fprint(w, fmt.Sprintf("%+v", words))
}

// this will add measuring execution time of handlerFunc passed as h argument
func measureExecTime(h http.HandlerFunc) http.HandlerFunc {
	// wrap handlerFunc to add some functionality
	return func(w http.ResponseWriter, r *http.Request) {
		// for measuring time we need start time and stop time
		// to get duration of action
		start := time.Now()
		// call h which is HandlerFunc
		h(w, r)
		// print to console execution duration with time.Since(Start))
		log.Println("handler execution time =>", time.Since(start))
	}
}

func main() {
	// register url and handlerFunc called for "/" url
	http.HandleFunc("/count", measureExecTime(countWords))
	// write that server starting in console
	fmt.Println("Starting Server ...")
	// start server and check for errors, maybe port is busy and it could
	// not start
	err := http.ListenAndServe("localhost:4000", nil)
	// fatal (stop program) on errors
	if err != nil {
		// fatal and log error
		log.Fatal("Server could not started", err)
	}

}
