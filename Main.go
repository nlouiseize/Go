package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	hostName        = "10.128.2.17"
	ClusterPort int = 3000
	namespace       = "test"
	setName         = "peeps"
	myKey           = "1"
	GoPort          = flag.Int("p", 8080, "server port")
	name            = "nelzir"
	age             = "30"
)

func main() {

	tmpl, err := template.ParseFiles("Aerospike-Inputs.html")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		fmt.Println("Server works!")

		hostName = r.FormValue("HostIP")
		namespace = r.FormValue("namespace")
		setName = r.FormValue("set")
		myKey = r.FormValue("PK")
		name = r.FormValue("name")
		age = r.FormValue("age")
		fmt.Println(name, age)
		
	})
	http.ListenAndServe(":8000", nil)
}
