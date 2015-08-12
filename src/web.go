package main

import (
	"contact"
	"encoding/json"
	"fmt"
	"net/http"
)

var contacts = []contact.Contact{}

func main() {

	//post /contacts
	//{"name" : "" , "tel" : "" , "email" : ""}
	http.HandleFunc("/contacts", addressBookHandler)
	// http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request){
	// 		w.Write([]byte("Hello Gopher"))
	// 	})
	//port 8000
	http.ListenAndServe(":8000", nil)
}

func addressBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var c contact.Contact
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			fmt.Println(err.Error())
		}
		contacts = append(contacts, c)
		fmt.Println(c)
	} else if r.Method == "GET" {
		name := r.URL.Query().Get("name")
		for _, c := range contacts {
			if c.Name == name {
				json.NewEncoder(w).Encode(c)
				return 
			}
		}

		w.WriteHeader(http.StatusNotFound)
	}
}
