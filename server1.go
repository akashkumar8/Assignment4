package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Move for the registration !") // write data to response
}

func reg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("reg.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("gender:", r.Form["gender"])
		fmt.Println("address:", r.Form["address"])
	}
}

func main() {
	http.HandleFunc("/", welcome) // setting router rule
	http.HandleFunc("/reg", reg)
	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}