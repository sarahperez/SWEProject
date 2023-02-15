//references:
//https://golang.ch/which-golang-router-to-use-for-what-cases/
// https://github.com/gorilla/mux
// https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// https://pkg.go.dev/net/http#ServeMux

// !!!!!!!!!!!!!!!!run go get github.com/rs/cors in terminal before running code!!!!!!!!!!!!!!!!

//after starting the server, open a new terminal and run these comands to test get http requests:
// curl.exe -v -X GET http://localhost:3000/home-page
// curl.exe -v -X GET http://localhost:3000/calendar
//curl.exe -v -X POST https://localhost:8888/login -H 'Content-Type: application/json' -d "@userInfo.json"

//run these commands to test post http requests:
//curl.exe -v -X POST http://localhost:3000/home-page -H 'Content-Type: application/json' -d "@userInfo.json"

// ctrl + c to terminate the server after using command go run .

package main

import (
	"encoding/json"
	"log"

	//web server router package- up to date (made by GO)
	"net/http"

	"github.com/jinzhu/gorm"

	//packages added from tutorial

	"main/database"
)

var DB *gorm.DB

// Main+functions were modified from: https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// the main function start the server
func main() {

	database.InitDatabase()
}

// practice struct to catch user data
type userinfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// this function will be called with the following URL: http://localhost:3000/home-page
// example code from https://golang.ch/which-golang-router-to-use-for-what-cases/ used as a reference
func goHome(w http.ResponseWriter, request *http.Request) {

	switch request.Method {
	//if the request is a GET
	case http.MethodGet:
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the home page, get request recieved"))

	case http.MethodPost:
		//if the request is a POST (incoming data)

		//reference for decoding (structure taken from example)- https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
		//set up a struct object to decode the json file into
		var info userinfo
		//decode the json file
		error := json.NewDecoder(request.Body).Decode(&info)
		if error != nil {
			//if statement to deal with decoder errors
			log.Println("decoding unsucessful", error)
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		//print the decoded info
		log.Println("decoded string:", info)
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//pass infromation back
		w.Write([]byte(info.Username))
		w.Write([]byte(info.Password))
	}
}

// this function will get called by the following URL: http://localhost:3000/calendar
// example code from https://golang.ch/which-golang-router-to-use-for-what-cases/ used as a reference
func displayCalendar(w http.ResponseWriter, request *http.Request) {

	switch request.Method {
	//if the request is a GET
	case http.MethodGet:
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the calendar page, get request recieved"))

	case http.MethodPost:
		//if the request is a POST (incoming data)
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the calendar page, post request recieved"))
	}
}

func defaultFunc() {

}

//old main function- this should be put back in main when it is established that the database is set up correctly
/*
	//database.InitDatabase()

	//put angular address here
	//host := ""

	//initalizing an HTTP request multiplexer- this can check to see if any of the incoming url match
	//those we load it with and then run the appropriate functions
	mux := mux.NewRouter()
	//here, we are telling the mux that if it gets passed the "/home-page" URL, to go to the goHome function
	mux.HandleFunc("/home-page", goHome)
	mux.HandleFunc("/calendar", displayCalendar)
	//mux.HandleFunc("/login", login).Methods("POST")
	//mux.HandleFunc("/register", register).Methods("POST")

	//trying to add in a handler for all cases where URL does NOT match one of the above linked to the mux
	//mux.PathPrefix("/").Handler()

	//set port (backend)
	const port = 3000
	//server will run on local host (your pc address)
	const server = "localhost"

	// the angular client loaded at http://localhost:4200 will make requests to the go server at http://localhost:3000
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
	//listen for requests sent to the server
	err := http.ListenAndServe(server+":"+strconv.Itoa(port), handler)

	//err := http.ListenAndServe(host, handler)
	log.Fatal(err)
*/
