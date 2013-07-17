// gorsi project main.go
package main

import (
	"flag"
	"fmt"
	"github.com/coocood/jas"
	"gorsi/model"
	"gorsi/routes"
	"log"
	"net/http"
	"utils"
)

var host_ip = flag.String("addr", ":8080", "Address to listen")

func main() {

	flag.Parse()

	//MODEL
	sporocilo.RegisterDb()
	sporocilo.UstvariTabeloSporocila()

	//REST
	router := jas.NewRouter(
		new(items.Items),  // last 25 items with :offset
		new(items.ItemId), // item operations [like, unlike, read, unread]
	)
	router.BasePath = "/v1/"
	router.EnableGzip = true
	fmt.Println(router.HandledPaths(true))
	http.Handle(router.BasePath, router)
	rest.RegisterWS()
	http.HandleFunc("/feed", rest.Atom)

	// Listen
	fmt.Println("Listening on", *host_ip)
	if err := http.ListenAndServe(*host_ip, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
