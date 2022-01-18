# autorequests

Creating API Wrappers have never been easier

## Features
- Write every requests no matter the method
- Can parse huge .har files in seconds
- Format code on save

## Command flags 

```
Usage of ./autorequests:
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
$ ./autorequests -i ./example/discord.com.har -exclude science -o ./example/discord.com.go
```
## Example
* See [example folder](https://github.com/bytixo/autorequests/tree/main/example) for output and input

```go

// Example of what an exported function look like

//DiscordCom_lgTeMaPEZQ make a POST request to https://discord.com/api/v9/auth/register
//with 24 Headers and 2 Cookies
func DiscordCom_lgTeMaPEZQ() {
	data := `{"fingerprint":"932960790959493181.999uKkMHr9uFYqRI9m_2Y8vPZIk","email":"harparser@gmail.com","username":"yothatsmeelol","password":"heyhowareyou667","invite":null,"consent":true,"date_of_birth":"1990-12-05","gift_code_sku_id":null,"captcha_key":"P0_eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwYXNza2V5Ijoic00zTWM1Zm5TSXR0d0k2TnZXL1d6c0ExZCtYWHlwNDJSL3grOG8zZ3BnR1pINE5keFRGdnpmQktkMVJSYjg0WnZwWjZXVlJ2MWxMSkVBcnlJK21UeHltR050UUtZVGNzVytqOW1kVUFKdzhjWXBHQitNVWNkUGYwU3Jxb0RCOEZmZXZmVnowOWZUQkwxSDNCV0Y0ZG1sUkNXYXZqai9YalYrbGJvN2dIeFpsQzRRNEo3Zk80Um5tWDB0Z1BiSWFyMmJoMTFiYlh1Rk1jR21hOElSemVxTGUzUGR3ZGp3TWU3RzA5R2l1b3NNWEtpeExmWDVQcUdnbXdobU9BZm9rUzFRcWUrTkRxT2VKTTZMbHdCdXFIczZhS3Q1bDdwcm9jcVRrdSs3QWdTd0VEVitJQlVlQ1IxUFNrNzA2Y0k4bDlSdlZXOEhXWkJ3ME1PQVF5eUZWQkFmMlFuT2E1c1RkemRXUTRpdDl2eG0rT0dOT2tpd256Y3BIdkJLM3JhM3NEaXNhWi9yN25naXczZ0plNHhFa0NObS9NNlA4WUNlaXE5Nmt0NTJPTENqV3VyK2FZNkR2T3NPbVpTSUpsb0F6a0VwK3A2SVhPekdoNVltb3NXZVB3VXcwM3k5RnJyL0hlREZSOVpXV2ROZFFYeXFhUFk5QkhiL1p4VW51cnhyRzVrUnVVVHhJMW1ENFJYT1ErcUFHSmFSU1dBZ3RBcDhiNk5TaWtJcldyeUxSZ3hiMUQ5TFNqd2d3UUdUN0R4Nm85Z3Z3bWtyVTFPeldqeEg1R2FhVEQycTdTN01VNWVWNlM3UXB3WURWbnJvd2JoMEFKbmh2Mmo3bEhvNVF6cGNEVFRLRzJBekhXWXM1bjRqM0RVWWtQd0IrU2dZRlhmT0k3SkgrdDJ5eUNXc3ZsUkZWRkJ2bktWSG83Rk9jQWVFTElGMTVjcW5hZUtKb05TN2hpTUFxR04wUVlMa2pzVTJnc1lmckkwNTNpQUllUlhNc1JqWW5QckhjWk5UZUt5anFDQ3RmS3lPR3ZkYkw3SWovMVpSWTJxdm9iTWZBWmNrNXkzdmdPNW56am5OVVFnY01aUVJCYjQwbURIMGxTMGJ5MTc2SE1pNnBRQ2owVEI4ZnJMaXg4eWljZ0doTzdTMlRLNkpBemNETjFxYzl3d2Y2ajRPWjRudUs5TG9DL3ArSVAzS0o1YVFGSWVyUWhKR05yNGd1TzFFZG1zUllHWVN2NWFNQXM2VWJpZDhKdjRPRjFkbFZKbVF1VmRsbnBrNHlvWUdFKzVzYmdRRzRybjVzUVdwN09SOUpvYUJSb084Q3p3a0lyM0kreldrQjdIVUZCeHJvdzY0WXNudWhKVHB3ejJGUXJGYVhKTWhEdCt6WHZQTXREaFY1M1YrMXlnZUU2TFNYblc0ekhDb1BtdWJFWHozSEhYdHlHZFg1Q1FUYkFacUZpdFZPdWovNCtJSE81c1VJaksybnFCOUdBYzhyK2Z5TCtzQTNZY2hMOXJWUTUvMlBEdFVCLzR2bkh5S3Yrdm5IcDVBdFRDN3AzVFpMcTRmU2FBNjhKcVBEb243b1VjTEhvZlFlZlI3S1JFMUE5WldiQXo5UnVCcnJabVpnVnV0cEVYdWFWbm9STGRhR053VVNPVTZ2MzN6Uk90ZkNlUlFZUnlyZkgvWDVwUU5RcTZOZFZYeTJzYmNtY2lWN2YyV051S21GY2VrMThZNTBWc2V1YUI3d3RyUnVyVWovcEtZRG1BYkY3N05KUG4wKzR3VkJqMHZ5dXMvTHQxZlpkRWl6emVXWDIvZU1WanhOU1FLWUtSendGanBkOHorUmJkbDRRK3YwVGwvUm8rdnY3S2RaRjhBM0VpQ1hicEJSQVA2b2h4TDM2blZtUEFSK0V2MzR3c2s4K2ZwVXp1UmJ6VHg2SEdUbk9oR283RDZvaWxtUzgwdXllYngrOWIzRDJ1Mjh6eVVKRHViNkp0M3k0YjhWRGN0SWZQc0NZeWtJV3lIdVZ6dHNoUEVnNDh0UGNqSmdRemlJNXZGM1hmTWJNUko1endkbjJiVlRxZk9NUjVnQkZhWDhCblRsK2RDNFNSUWM0VEYzd3lkbGFleUMxRDY1K1VoQWZCeVBrTTk5UDRCNmJUeFZZRCt4TWF1enFDQThuRndnUko1N2JtNnkwanVSOEFabzUwdHRZSXlxV0NQTWhJOW5BZnBUbUJZcE1MUGxFY1dUSCtkdmlhaXNHTkxRMlpmSUUvSTlmMmVEOTRHaGNtTHA5TlRiZlIzYUhlM2Exc21wNUU0MEc0OWx5VzYzbjBzOFRMbUFFSk9sQitOSDl1VW9FSm4zWDNCM1g3dWZtdVcyenpGUGYvRUtaVzNYVFhhM2ZJanBRbTlsNGk3Q2IxM1NBZXVtQzdhVFhwTHE2MDJDSmRtK0MiLCJleHAiOjE2NDI1MDU3NDksInNoYXJkX2lkIjo4MjA3ODYwODYsInBkIjowfQ.NKfvK9_MS1yKBp_P6t9uwF-nTEAiafU8zNc_j6HTPZU"}`

	req, err := http.NewRequest("POST", "https://discord.com/api/v9/auth/register", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Content-Length":     []string{fmt.Sprint(len(data))},
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



