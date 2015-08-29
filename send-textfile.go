// create utility which will
// read text file and send it 
// to count word server
// after receiving result it
// should print it out
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	fatalOnErr(err)

	res, err := http.Post("http://localhost:4000/count", "text/txt", f)
	fatalOnErr(err)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fatalOnErr(err)
	fmt.Println("response", string(body))
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
