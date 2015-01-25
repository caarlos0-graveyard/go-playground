package main

import (
	"flag"
	"github.com/caarlos0/go-playground/essential-go/pages"
	"log"
	"os"
)

func main() {
	config := flag.String("config", "config.json", "The config.json file to use")
	layout := flag.String("layout", "layout.html", "The layout.html file to use")
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	filename := flag.Args()[0]
	p, err := pages.NewPage(filename, *config)
	if err != nil {
		log.Fatalln(err)
	}
	err = p.Render(*layout, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
