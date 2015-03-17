package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strconv"
	"time"
	"os"
)

//ROUTING

func main() {

	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/count", count).Methods("GET")
	mainRouter.HandleFunc("/add", add).Methods("POST")

	http.Handle("/", mainRouter)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)	}
}

func add(w http.ResponseWriter, r *http.Request) {

	urlValues := r.URL.Query()
	string := urlValues.Get("string")

	query := bson.M{
		"s": string,
	}

	s, err := mgo.DialWithTimeout("127.0.0.1:27017", 100*time.Millisecond)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		c := s.DB("daffe").C("dummy")

		_ = c.Insert(query)

		w.Write([]byte("OK!!!"))
		s.Close()
	}

}

func count(w http.ResponseWriter, r *http.Request) {
  addr := os.GetEnv('MONGO_PORT_27017_TCP_ADDR')
  port := os.GetEnv('MONGO_PORT_27017_TCP_PORT')
  addrport := fmt.Sprintf('%s:%s', addr, port)
	s, err := mgo.DialWithTimeout(addrport, 100*time.Millisecond)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s - %s", err, addrport)))
	} else {
		c := s.DB("daffe").C("dummy")

		ct, _ := c.Count()

		w.Write([]byte(strconv.Itoa(ct)))
		s.Close()
	}

}
