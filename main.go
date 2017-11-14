package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"Cloudgo/service"
)

const (
	//defult port num 
	PORT string = "8080"
)

func main() {
	//get port
	port := os.Getenv("PORT")
	// if not given,port==8080
	if len(port) == 0 {
		port = PORT
	}
	Port := flag.StringP("port", "p", PORT, "PORT for hpptd listening")
	flag.Parse()
	if len(*Port) != 0 {
		port = *Port
	}
	service.NewServer(port)
}
