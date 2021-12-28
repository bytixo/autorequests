# HAR-Parser
Parse .har files to create Go http requests


## Usage

```sh
$ ./HAR-Parser myfile.har
```

## Example
* See [example folder](https://github.com/bytixo/HAR-Parser/tree/main/example) for output and input

```go

// Example of what an exported function look like

func CanaryDiscordCom_bZRjxAwnwe() {

	// Type: application/json

	data := `{"events":[{"type":"network_action_user_register","properties":{"client_track_timestamp":1640551449221,"status_code":400,"url":"/auth/register","request_method":"post","invite_code":null,"client_performance_cpu":3.6445815497355323,"client_performance_memory":220232,"cpu_core_count":12,"accessibility_support_enabled":false,"accessibility_features":256,"rendered_locale":"fr","client_uuid":"KTBCsuxs0Qys6N7DpDR8+H0BAAAMAAAA","client_send_timestamp":1640551449231}},{"type":"impression_user_registration","properties":{"client_track_timestamp":1640551449225,"impression_type":"view","impression_group":"user_registration_flow","step":"captcha","location_section":"impression_user_registration","client_performance_cpu":3.6445815497355323,"client_performance_memory":220232,"cpu_core_count":12,"accessibility_support_enabled":false,"accessibility_features":256,"rendered_locale":"fr","client_uuid":"KTBCsuxs0Qys6N7DpDR8+H0BAAANAAAA","client_send_timestamp":1640551449231}}]}`

	req, err := http.NewRequest("POST", "https://canary.discord.com/api/v9/science", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Content-Length":	[]string{`972`},
		"X-Super-Properties":	[]string{`eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJjYW5hcnkiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC40MiIsIm9zX3ZlcnNpb24iOiIxMC4wLjE5MDQyIiwib3NfYXJjaCI6Ing2NCIsInN5c3RlbV9sb2NhbGUiOiJmciIsImNsaWVudF9idWlsZF9udW1iZXIiOjEwODk5MiwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Fingerprint":	[]string{`923639162446426153.-Ru9bDrWlyojIhYzD9gWgz-7-SI`},
		"X-Debug-Options":	[]string{`bugReporterEnabled`},
		"Accept-Language":	[]string{`fr,fr-FR;q=0.9,en-US;q=0.8`},
		"Authorization":	[]string{`undefined`},
		"Content-Type":	[]string{`application/json`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.42 Chrome/91.0.4472.164 Electron/13.4.0 Safari/537.36`},
		"X-Discord-Locale":	[]string{`fr`},
		"Accept":	[]string{`*/*`},
		"Origin":	[]string{`https://canary.discord.com`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://canary.discord.com/register`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Cookie":	[]string{`__dcfduid=711333b02ea811ec9a6703a73b1dd318; __sdcfduid=711333b12ea811ec9a6703a73b1dd31846ee9709ba9b89980eea8fdb13a16e5db35956253dc2f55b0a3dc240cc4703c8`},
		
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
```

### Video

https://user-images.githubusercontent.com/69140663/147421063-0651308c-439b-438f-94b9-5029df1a032c.mp4

