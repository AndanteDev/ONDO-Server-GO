package main

import (
	"log"
	"ondo/server/go/handler"
	"ondo/server/go/info"
)

func main() {

	r := handler.MakeHandler()

	log.Fatal(r.Run(info.Port))
	// log.Fatal(r.RunTLS(info.Port, "server.crt", "server.key"))

}
