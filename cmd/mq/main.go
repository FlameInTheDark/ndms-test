package main

import (
	"flag"
	"log"

	"github.com/FlameInTheDark/ndms-test/internal/api"
)

var url = flag.String("url", "localhost:8080", "host url")

func main() {
	flag.Parse()

	app := api.NewAPI()
	log.Fatal(app.Serve(*url))
}
