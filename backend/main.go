//references - https://golang.ch/which-golang-router-to-use-for-what-cases/
// https://github.com/gorilla/mux
// https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// https://pkg.go.dev/net/http#ServeMux

// !!!!!!!!!!!!!!!!run go get github.com/rs/cors in terminal before running code!!!!!!!!!!!!!!!!
//ctrl + c to terminate the server after using command go run .

package main

import (
	"log"
	//web server router package- up to date (made by GO)
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

// the main function start the server
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-page-data", example)

	//set port (backend)
	const port = 3000
	//server will run on local host (your pc address)
	const server = "localhost"

	// the angular client will be loaded at http://localhost:4200 and make requests to the go server at http://localhost:3000
	// since these addresses are different we need to set a CORS policy to allow requests from our client
	c := cors.New(cors.Options{
		//tell computer that it can accept requests the origin of the angular app
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	//handler is assigned the "function" that will call getPageData if it passess the CORS and matches the path
	handler := c.Handler(mux)

	//log.Printf shows date and time- could also just use Printf, but log better practice
	log.Printf("starting server on http://%s:%d", server, port)
	//start the web server
	//listen for requests sent to ("on" proper terminology) 3000
	err := http.ListenAndServe(server+":"+strconv.Itoa(port), handler)
	//if something does not work, (exit status 1) ie. if someone tries to use the same port
	log.Fatal(err)
}

// handler function for requests to http://server:port/get-page-data
// more info on request types https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods
func example(w http.ResponseWriter, r *http.Request) {

	log.Println("getPageData: %s", r.URL.Path)

	switch r.Method {
	//if the request is a GET
	case http.MethodGet:
		//tell in json format
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("{\"title\":\"Test Sucessful\"}"))
	}
}
