package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	now := time.Now()

	/*
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("cannot get hostname panic !!!!")
		//panic(err)
	} else {
		fmt.Println(host)
	}
	*/
	host := os.Getenv("HOSTNAME")
	if host != "" {

		host = "na"

	}

	//var name string
	//name = "oliver go"
	//name2 := "I love Vanessa"
	fmt.Fprintf(w, "Hello Go test - " + now.String() + ": " +host) // send data to client side
	//fmt.Fprintf(w, "Hello " + name + " " + name2 ) // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
