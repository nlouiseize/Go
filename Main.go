package main

import (
	"fmt"
	aero "github.com/aerospike/aerospike-client-go"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	hostName        = "10.128.2.24"
	port        int = 3000
	namespace       = "test"
	setName         = "peeps"
	myKey       int = 1
	User            = "admin"
	Password        = "admin123"
	ClusterName     = "aeroclustersrc"
)

func main() {

	// define a client to connect to

	var clientPolicy = aero.NewClientPolicy()
	clientPolicy.User = User
	clientPolicy.Password = Password
	//clientPolicy.ClusterName = ClusterName

	client, err := aero.NewClientWithPolicy(clientPolicy, hostName, port)
	panicOnError(err)

	//client, err := aero.NewClient(hostName, port)
	//panicOnError(err)

	// Create new write policy
	policy := aero.NewWritePolicy(0, 0)
	policy.SendKey = true

	key, err := aero.NewKey(namespace, setName, myKey)
	panicOnError(err)

	// define some bins with data
	bins := aero.BinMap{
		//	"PK":   myKey,
		"name": "nelzir",
		"age":  29,
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

	// // delete the key, and check if key exists
	// existed, err := client.Delete(nil, key)
	// panicOnError(err)
	// fmt.Printf("Record existed before delete? %v\n", existed)
}
