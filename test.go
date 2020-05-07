package main

import (
	"awdb/awdb-golang"
	"fmt"
	"log"
	"net"
)

func main() {
	db, err := awdb.Open("GeoIP2-City.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP("1.192.90.1")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", record)

}