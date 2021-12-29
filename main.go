package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type Request struct {
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	HTTPVersion string    `json:"httpVersion"`
	Headers     []Headers `json:"headers"`
	Cookies     []Cookies `json:"cookies"`
	PostData    PostData  `json:"postData,omitempty"`
}
type Headers struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Cookies struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	Expires  time.Time `json:"expires"`
	HTTPOnly bool      `json:"httpOnly"`
	Secure   bool      `json:"secure"`
	SameSite string    `json:"sameSite,omitempty"`
}
type PostData struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}

func main() {
	var filename string
	filename = os.Args[1]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	j := string(b)
	value := gjson.Get(j, "log.entries")

	for _, entry := range value.Array() {
		if entry.Get("_resourceType").String() == "fetch" || entry.Get("_resourceType").String() == "xhr" {

			var request Request
			err := json.Unmarshal([]byte(entry.Get("request").String()), &request)

			if err != nil {
				fmt.Println("Error while unmarshaling:", err)
			}

			w := parseRequest(request)

			f := strings.TrimSuffix(filename, ".har")
			f = strings.Title(strings.ReplaceAll(f, ".", " "))
			f = strings.ReplaceAll(f, " ", "") + ".go"

			writeToFile(f, w+"\n")
		}
	}

}

func parseRequest(request Request) string {

	parsedURL, _ := url.Parse(request.URL)
	r := strings.NewReplacer(".", " ", "/", " ")
	funcName := strings.Title(r.Replace(parsedURL.Host))
	funcName = strings.ReplaceAll(funcName, " ", "") + "_" + RandStringBytes(10)

	encodingType := request.PostData.MimeType
	data := request.PostData.Text
	method := request.Method
	URL := request.URL
	info := fmt.Sprintf(`"%s" request @ %s`, method, URL)

	l := "`"
	var headers string
	for _, h := range request.Headers {
		if strings.HasPrefix(h.Name, ":") {
			continue
		}
		headers += fmt.Sprintf(`"%s":	[]string{%s},
		`, strings.Title(h.Name), l+h.Value+l)
	}

	final := fmt.Sprintf(
		`
func %s() {


	/* 

		Info :%s

		Type: %s

	*/ 

	data := %s

	req, err := http.NewRequest("%s", "%s", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		%s
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		`, funcName, info, encodingType, l+data+l, method, URL, headers)

	return final
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func writeToFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
}
