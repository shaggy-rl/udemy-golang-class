package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
_ is the blank identifier. In this case http.Get returns a result and error
since we are not going to use error in our code we replace the variable def
with a _ to tell the compiler that this is intended and to throw it away
*/
func main() {
	res, _ := http.Get("http://mcleads.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
