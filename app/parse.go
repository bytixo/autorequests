package app

import (
	"fmt"
	"go/format"
	"math/rand"
	"net/url"
	"strings"
)

var (
	l = "`"
)

func (c *Config) ParseRequests(requests []Request) (string, error) {
	var result string

	result += "package app\n\nimport (\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"net/http\"\n\t\"strings\"\n\t\"time\"\n)\n"

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
		data := request.PostData.Text
		method := request.Method
		URL := request.URL

		comment := fmt.Sprintf("//%s make a %s request to %s\n//content type: %s", funcName, method, URL, encodingType)

		// headers
		var headers string
		for _, h := range request.Headers {
			// skip :authority :scheme etc
			if strings.HasPrefix(h.Name, ":") {
				continue
			}
			headers += fmt.Sprintf(`"%s":	[]string{%s},
		`, strings.Title(h.Name), l+h.Value+l)
		}

		result = fmt.Sprintf(
			`
%s
func %s() {
	data := %s

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
		`, comment, funcName, l+data+l, method, URL, headers)

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
