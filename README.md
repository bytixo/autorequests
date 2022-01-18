# HAR-Parser

Creating API Wrappers have never been easier

## Features
- Write every requests no matter the method
- Can parse huge .har files in seconds
- Format code on save

## Command flags 

```
Usage of ./HAR-Parser:
  -exclude string
        comma separated list of URL patterns you would like to exclude. Providing "-exclude facebook,science" will remove every URL that contains "facebook" AND/OR "science"
  -i string
        input to your .har file
  -o string
        where it should write the result, by default it's almost the same as your original .har file
```
## Example Usage

this will skip URL's that contains "science"
```sh
$ ./HAR-Parser -i ./example/discord.com.har -exclude science -o ./example/discord.com.go
```
## Example
* See [example folder](https://github.com/bytixo/HAR-Parser/tree/main/example) for output and input

```go

// Example of what an exported function look like

//DiscordCom_tcuAxhxKQF make a POST request to https://discord.com/api/v9/auth/register
//content type: application/json
func DiscordCom_tcuAxhxKQF() {
	data := `{"fingerprint":"932960790959493181.999uKkMHr9uFYqRI9m_2Y8vPZIk","email":"harparser@gmail.com","username":"yothatsmeelol","password":"heyhowareyou667","invite":null,"consent":true,"date_of_birth":"1990-12-05","gift_code_sku_id":null,"captcha_key":null}`

	req, err := http.NewRequest("POST", "https://discord.com/api/v9/auth/register", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Content-Length":     []string{`251`},
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Fingerprint":      []string{`932960790959493181.999uKkMHr9uFYqRI9m_2Y8vPZIk`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`undefined`},
		"Content-Type":       []string{`application/json`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://discord.com`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/register`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
		"Cookie":             []string{`__dcfduid=659b7120785211ecaed6a5158183df4c; __sdcfduid=659b7121785211ecaed6a5158183df4c37874ed17fa4c1a0d743a29588308044f9229446a217f31c602a11fa49438ee4`},
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
```



