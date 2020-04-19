package birthday

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, `{"msg": "Happy Birthday Service!"}`)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"msg": "Hello People!"}`)
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
	w.WriteHeader(http.StatusOK)

	name := html.EscapeString(r.URL.Path)
	name = name[7:]
	var helloMan HelloMan
	helloMan.Name = name

	// man := RepoFindMan(name)
	if FindCollection(helloMan) {
		msg := WhenIsBirthday(helloMan.DateOfBirth, name)
		//fmt.Println(msg) //log to console
		jsonMap := map[string]string{"msg": msg}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, `{"msg":"Unfortunatelly, name '%s' is not in the database."}`, name)
	}
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
	}

	name := html.EscapeString(r.URL.Path)
	name = name[7:]
	helloMan.Name = name

	//man := RepoFindMan(name)
	if FindCollection(helloMan) {
		//RepoUpdateMan(name, helloMan.DateOfBirth)
		UpdateCollection(helloMan, helloMan)
	} else {
		//RepoCreateMan(name, helloMan.DateOfBirth)
		SaveCollection(helloMan)
	}

	w.WriteHeader(http.StatusNoContent)
}
