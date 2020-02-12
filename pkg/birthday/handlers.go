package birthday

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy Birthday Service")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"msg": Hello People}`)
	// prints list of people
	/*
	fmt.Fprintln(w, "Hello People function.\nData already saved:")
 	if err := json.NewEncoder(w).Encode(helloPeople); err != nil {
		panic(err)
	}
	*/
}

func HelloSomebody(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	name := html.EscapeString(r.URL.Path)
	name = name[7:]

	man := RepoFindMan(name)
	if (man != (HelloMan{}) ) {
		msg := WhenIsBirthday(man.DateOfBirth, name)
		fmt.Println(msg) //log to console
		jsonMap := map[string]string{"message": msg}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, "Unfortunatelly, the name '%s' is not in the database.", name)
	}
	//w.WriteHeader(http.StatusOK)
}

func SaveSmbsName(w http.ResponseWriter, r *http.Request) {
	var helloMan HelloMan
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &helloMan); err != nil {
		w.Header().Set("content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		panic(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	name := html.EscapeString(r.URL.Path)
	name = name[7:]

	man := RepoFindMan(name)
	if (man != (HelloMan{}) ) {
		RepoUpdateMan(helloMan.DateOfBirth, name)
	} else {
		RepoCreateMan(helloMan, name)
	}

	w.WriteHeader(http.StatusNoContent)
	/*
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	*/
}

var mu sync.Mutex
var count int

func Counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Number of request equals: %d\n", count)
	mu.Unlock()
}

func Parser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
