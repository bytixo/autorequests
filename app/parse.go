package app

import (
	"fmt"
	"go/format"
	"log"
	"math/rand"
	"net/url"
	"strings"
)

var (
	l          = "`"
	tabulation = "\n\t"
)

func (c *Config) ParseRequests(requests []Request) (string, error) {
	var result string

	result += "package main\n\nimport (\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"net/http\"\n\t\"strings\"\n\t\"time\"\n\t\"net/url\"\n)\n"

	for _, request := range requests {
		result += parseRequest(request, c.Exclude)
	}
	content, err := format.Source([]byte(result))
	if err != nil {
		return "", fmt.Errorf("go/format error: %s", err)
	}
	return string(content), nil
}

func parseRequest(request Request, excludes []string) string {
	var result string
	// Might be a better way of doing this but idk
	for _, exclude := range excludes {
		if strings.Contains(request.URL, exclude) {
			continue
		}
		// Create our func name TODO: make it look better
		pURL, _ := url.Parse(request.URL)
		r := strings.NewReplacer(".", " ", "/", " ")
		funcName := strings.Title(r.Replace(pURL.Host))
		funcName = strings.ReplaceAll(funcName, " ", "") + "_" + RandStringBytes(10)

		encodingType := request.PostData.MimeType

		var data string
		switch {

		// support for url query params
		case encodingType == "application/x-www-form-urlencoded":
			values, err := url.ParseQuery(request.PostData.Text)
			if err != nil {
				log.Println("can't parse URL query, skipping ...")
				data = fmt.Sprintf(`data := %s`, l+request.PostData.Text+l)
			}
			data += "form := url.Values{}" + tabulation
			for k, v := range values {
				for _, value := range v {
					data += fmt.Sprintf(`form.Add("%s",%s)`, k, l+value+l) + tabulation
				}
			}
			data += "data := form.Encode()" + tabulation
		default:
			data += fmt.Sprintf("data := `%s`", request.PostData.Text)
		}
		method := request.Method
		URL := request.URL

		comment := fmt.Sprintf("//%s make a %s request to %s\n//with %d Headers and %d Cookies", funcName, method, URL, len(request.Headers), len(request.Cookies))

		// headers
		var headers string
		for _, h := range request.Headers {

			switch {
			// skip :authority :scheme etc
			case strings.HasPrefix(h.Name, ":"):
				continue
			// we want to set a special value here
			case strings.Contains(h.Name, "-length"):
				headers += fmt.Sprintf(`"%s":	[]string{%s},
				`, strings.Title(h.Name), "fmt.Sprint(len(data))")
			default:
				headers += fmt.Sprintf(`"%s":	[]string{%s},
				`, strings.Title(h.Name), l+h.Value+l)
			}

		}

		result = fmt.Sprintf(
			`
%s
func %s() {
	%s

	req, err := http.NewRequest("%s", "%s", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		%s
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		// handle error ...
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error ...
	}
	// do something ...
}
		`, comment, funcName, data, method, URL, headers)
	}
	return result

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
