package main

import (
	"flag"
	"fmt"
	aero "github.com/aerospike/aerospike-client-go"
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
	bin1name        = "name"
	bin1value       = "nelzir"
	bin2name        = "age"
	bin2value       = "30"
)

func main() {

	// An HTML template
	const tmpl = "<div>\n    {{if .Success}}\n    <div>\n        <p>Record is successfully created!</p>\n    </div>\n    {{else}}\n    <div>\n        <h1>Aerospike Inputs</h1>\n    </div>\n    <div></div>\n    <form action=\"/aerospike-inputs.html\" method=\"POST\">\n        <input\n                type=\"text\" placeholder=\" host IP\" name=\"HostIP\"\n        />\n        <input\n                type=\"text\" placeholder=\" namespace\" name=\"namespace\"\n        />\n        <input\n                type=\"text\" placeholder=\" set\" name=\"set\"\n        />\n        <input\n                type=\"text\" placeholder=\" PK\" name=\"PK\"\n        />\n        <input\n                type=\"text\" placeholder=\" bin1 name\" name=\"bin1name\"\n        />\n        <input\n                type=\"text\" placeholder=\" bin1 value\" name=\"bin1value\"\n        />\n        <input\n                type=\"text\" placeholder=\" bin2 name\" name=\"bin2name\"\n        />\n        <input\n                type=\"text\" placeholder=\" bin2 value\" name=\"bin2value\"\n        />\n        <input  type=\"submit\" value=\"submit\" />\n    </form>\n    {{end}}\n</div>"

	t, err := template.New("Aerospike-Inputs").Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Execute(w, nil)
			return
		}
		fmt.Println("Server works!")
		err := t.Execute(w, struct{ Success bool }{true})
		if err != nil {
			return
		}

		hostName = r.FormValue("HostIP")
		namespace = r.FormValue("namespace")
		setName = r.FormValue("set")
		myKey = r.FormValue("PK")
		bin1name = r.FormValue("bin1name")
		bin1value = r.FormValue("bin1value")
		bin2name = r.FormValue("bin2name")
		bin2value = r.FormValue("bin2value")

		client, err := aero.NewClient(hostName, ClusterPort)
		panicOnError(err)

		// Create new write policy
		policy := aero.NewWritePolicy(0, 0)
		policy.SendKey = true

		key, err := aero.NewKey(namespace, setName, myKey)
		panicOnError(err)

		// define some bins with data
		bins := aero.BinMap{
			//	"PK":   myKey,
			bin1name: bin1value,
			bin2name: bin2value,
		}

		// write the bins
		err = client.Put(policy, key, bins)
		panicOnError(err)

		// read it back!
		rec, err := client.Get(nil, key)
		panicOnError(err)
		//fmt.Println(rec)
		fmt.Printf("Record: %v", rec.Bins)

		client.Close()

	})
	http.ListenAndServe(":8080", nil)

}
