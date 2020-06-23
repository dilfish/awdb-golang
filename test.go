package main

import (
	"./awdb-golang"
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

	ip := net.ParseIP("166.111.4.100")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", record)

}
