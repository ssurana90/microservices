package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Handlers are objects that “handle” HTTP request
	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		//log.Println("hello world")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Data is %s\n", data)
	})
	//The http.ListenAndServe function internally creates a tcp listener on address addr using net.Listen function
	//which returns a net.Listener and uses it with http.Serve function to listen to incoming connections using handler
	http.ListenAndServe(":9090", nil)
}
