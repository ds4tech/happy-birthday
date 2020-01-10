package birthday

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"sync"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy Birthday service")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello People function.\nData saved already in JSON:\n")
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helloPeople); err != nil {
		panic(err)
	}
}

func HelloSomebody(w http.ResponseWriter, r *http.Request) {
	name := html.EscapeString(r.URL.Path)
	name = name[7:]

	man := RepoFindMan(name)
	if (man != (HelloMan{}) ) {
		dateOfToday := time.Now().Format(timeISO)

		if (dateOfToday == man.DateOfBirth) {
			fmt.Printf("Hello, %s! Happy birthday!\n", name)
			fmt.Fprintf(w, "Hello, %s! Happy birthday!", name)
		} else if ( IsBirthdayIn5Days(man.DateOfBirth) ) {
			fmt.Printf("Hello, %s! Your birthday is in 5 days!\n", name)
			fmt.Fprintf(w, "Hello, %s! Your birthday is in 5 days!", name)
		} else {
				fmt.Printf("Hello, %s!\n", name)
				fmt.Fprintf(w, "Hello, %s!", name)
		}

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
		RepoUpdateMan(helloMan, name)
	} else {
		RepoCreateMan(helloMan, name)
	}

	w.Header().Set("content-Type", "application/json; charset=UTF-8")
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
