package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
)

//ROUTING

func main() {

	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/count", count).Methods("GET")
	mainRouter.HandleFunc("/add",add).Methods("POST")
	
	http.Handle("/", mainRouter)
	
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
	      panic(err)
	}
}


func add(w http.ResponseWriter, r *http.Request) {

	urlValues := r.URL.Query()
	string := urlValues.Get("string")

	query := bson.M{
		"s": string,
	}
	
	s, err := mgo.Dial("localhost:27017")
	defer s.Close()
	
	if err != nil {
		panic(err)
	}

	c := s.DB("daffe").C("dummy")

    err = c.Insert(query)

    w.Write([]byte("OK"))
	
}


func count(w http.ResponseWriter, r *http.Request) {
	
	s, err := mgo.Dial("localhost:27017")
	defer s.Close()
	
	if err != nil {
		panic(err)
	}

	c := s.DB("daffe").C("dummy")

    ct, err := c.Count()
	
    w.Write([]byte(strconv.Itoa(ct)))
	
}
