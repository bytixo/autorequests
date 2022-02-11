package main

import (
	"math/rand"
	"time"
	"flag"
	"github.com/bytixo/autorequests/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	var exclude string
	var output string
	var filename string

	flag.StringVar(&filename, "i", "", "input to your .har file")
	flag.StringVar(&exclude, "exclude", "", `comma separated list of URL patterns you would like to exclude. Providing "-exclude facebook,science" will remove every URL that contains "facebook" AND/OR "science"`)
	flag.StringVar(&output, "o", "", "where it should write the result, by default it's almost the same as your original .har file")
	flag.Parse()

	app.Run(filename, exclude, output)

}
