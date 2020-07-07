package main

import (
	"./awdb-golang"
	"fmt"
	"log"
	"net"
)

func main() {
	db, err := awdb.Open("IP_city_single_BD09_WGS84_ipv6_awdb.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP("2400:da00::dbf:0:100")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", record)

}
