package main

import (
	"flag"
	"log"
	"os"
)

var (
	help            = flag.Bool("h", false, "help")
	updateDatabase  = flag.Bool("u", false, "update database")
	degradeDatabase = flag.String("d", "", "degrade database")
)

type Ftype func() error

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *updateDatabase {
		log.Println("update database successful")
	} else if *degradeDatabase != "" {
		log.Println("degrade database successful")
	} else {
		log.Println("Capricorn server, listen")
	}
}
