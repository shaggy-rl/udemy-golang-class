package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Since we are using error checking we want to save the err response from
http.Get, so we don't use the _ (blank identifier)
*/
func main() {
	res, err := http.Get("http://www.shaggy.pw")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
