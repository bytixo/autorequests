package app

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func Run(filename, exclude, output string) {
	conf := &Config{}
	if filename == "" {
		panic("you need to provide a file")
	}
	var excludes []string

	if exclude != "" {
		excludes = strings.Split(exclude, ",")
		conf.Exclude = excludes
	}

	requests, err := getRequests(filename)
	if err != nil {
		log.Println("error while parsing .har file:", err)
	}

	result, err := conf.ParseRequests(requests)
	if err != nil {
		log.Println("error while parsing requests", err)
	}

	switch {
	case output != "":
		err := writeToFile(output, result)
		if err != nil {
			log.Println("error writing result:", err)
		}
	default:
		var out string
		out = strings.TrimSuffix(filename, ".har")
		out = out + "_" + time.Now().String() + ".go"

		err := writeToFile(out, result)
		if err != nil {
			log.Println("error writing result:", err)
		}
	}

	log.Println("Finished parsing", filename)

}

func getRequests(filename string) ([]Request, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var requests []Request
	value := gjson.Get(string(b), "log.entries")

	for _, entry := range value.Array() {
		if entry.Get("_resourceType").String() == "fetch" || entry.Get("_resourceType").String() == "xhr" {

			var request Request
			err := json.Unmarshal([]byte(entry.Get("request").String()), &request)

			if err != nil {
				return nil, err
			}

			requests = append(requests, request)

		}
	}

	return requests, nil
}
func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
