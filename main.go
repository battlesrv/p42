package main

import (
	"fmt"
	"log"
	"os"

	"github.com/battlesrv/p42/db"
	"github.com/battlesrv/p42/routers"

	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	app := cli.NewApp()
	app.Author = "Konstantin Kruglov"
	app.Email = "kruglovk@gmail.comm"
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Action:  server,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "port",
					Usage: "bind to port",
					Value: 5000,
				},
				cli.StringFlag{
					Name:  "dbaddr",
					Usage: "address of DB tokens",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "dbport",
					Usage: "port of DB tokens",
					Value: 3000,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func server(c *cli.Context) {
	db.NewConn(c.String("dbaddr"), c.Int("dbport"))

	srv := routers.Init()
	srv.Run(fmt.Sprintf(":%d", c.Uint("port")))
}
