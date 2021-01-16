package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:desc`
	Content string `json:content`
}

// Declaration of empty Articl array
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test title", Desc: "Test Desc", Content: "Test content"},
	}

	fmt.Println("Endpoint Hit: all Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func allPostArticles(w http.ResponseWriter, r *http.Request) {
	// Response send send
	fmt.Fprintf(w, "allPostArticles Endpoint")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	// Response send send
	fmt.Fprintf(w, "HomePager EndPoint Hit")
}

func handleRequest() {

	myRoute := mux.NewRouter().StrictSlash(true)
	// Declaration of api route
	// homePage is func homePage which is called in handle func
	myRoute.HandleFunc("/", homePage)
	myRoute.HandleFunc("/articles", allArticles).Methods("GET")
	myRoute.HandleFunc("/allPostArticles", allPostArticles).Methods("POST")
	// Declaration of server and port
	log.Fatal(http.ListenAndServe(":8081", myRoute))
}

func main() {
	// call the server to initialize
	handleRequest()
}
