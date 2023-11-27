package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Database username")
	passsword := flag.String("password", "root", "Database password")
	host := flag.String("host", "127.0.0.1", "Database host")
	port := flag.String("port", "3306", "Database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *passsword)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	// run
	// go run flag.go -username=aditya -password="rahasia banget" -host=localhost -port=5432
}
