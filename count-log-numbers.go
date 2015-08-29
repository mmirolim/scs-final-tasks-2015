// create web server with two handlers
// one will generate logs directory and write 10 log files with
// random numbers from 1 to 1000 per line
// second handler will count sum of all numbers in all log files
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// generate 10 log files in logs directory
// each file contains random numbers per line from 1 to 1000
func generateLogFiles(w http.ResponseWriter, r *http.Request) {
	// print to console
	log.Println("generating log files")
	// create directory for logs
	err := os.Mkdir("logs", os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		log.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}
	// create 10 files with 1000 numbers per line
	for n := 0; n < 10; n++ {
		f, err := os.Create("logs/log" + strconv.Itoa(n) + ".txt")
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
		}
		data := ""
		for i := 0; i < 1000; i++ {
			data += strconv.Itoa(rand.Intn(1000)) + "\n"
		}
		f.Write([]byte(data))
	}
	fmt.Fprint(w, "done")
}

// count sum of all numbers in all log files in logs directory
// and show result
func count(w http.ResponseWriter, r *http.Request) {
	// print to console
	log.Println("start counting numbers in logs")
	// read directory
	dir, err := os.Open("logs")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}
	// read all files in directory
	fls, err := dir.Readdir(-1)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}
	sum := 0
	for _, f := range fls {
		// read whole files one by one
		dat, err := ioutil.ReadFile("logs/" + f.Name())
		// check for err
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
			return
		}
		// convert to string then split by \n then range
		// and convert ty int to sum it up
		for _, v := range strings.Split(string(dat), "\n") {
			d, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			sum += d
		}

	}
	// convert sum to string
	fmt.Fprint(w, strconv.Itoa(sum))
}

func main() {
	// register url and handlerFunc called for "/" url
	http.HandleFunc("/gen", generateLogFiles)
	http.HandleFunc("/count", count)

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
