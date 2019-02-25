package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/battlesrv/p42/db"
	"github.com/battlesrv/p42/routers"
)

var (
	port   = flag.Int("port", 5000, "bind to port")
	dbhost = flag.String("dbhost", "127.0.0.1", "address of DB tokens")
	dbport = flag.Int("dbport", 3000, "port of DB tokens")
)

func main() {
	server := routers.Init()
	server.Run(fmt.Sprintf(":%d", *port))
}

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	flag.Parse()

	db.NewConn(*dbhost, *dbport)
}
