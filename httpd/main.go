package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	//log "github.com/cihub/seelog"
)

func check(e error) bool {
	if e != nil {
		log.Fatal(e)
		return false
	} else {
		return true
	}

}

func test(filename string) {
	d1 := []byte("hello\ngo\n")
	log.Printf("writeing to file %v", filename)
	err := ioutil.WriteFile(filename, d1, 0644)
	check(err)
}

func test2(filename string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()
}

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

	/*
		if host != "" {

			host = "na"

		}
	*/

	//var name string
	//name = "oliver go"
	//name2 := "I love Vanessa"
	fmt.Fprintf(w, "<html><body><h1>Hello Go test </h1>- "+now.String()+": "+host + "<body></html>")  // send data to client side
	log.Println("Hello Go test - " + now.String() + ": " + host)
	//fmt.Fprintf(w, "Hello " + name + " " + name2 ) // send data to client side

}

func main() {
	log.Println("write file")

	test("/tmp/filename.txt")
	test2("/tmp/filename2.txt")

	log.Println("Starting Server !")
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
