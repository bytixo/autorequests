package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//autorequests_XVlBzgbaiCMRAjW make a GET request to https://discord.com/api/v9/experiments
//with 21 Headers and 2 Cookies
func autorequests_XVlBzgbaiCMRAjW() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/experiments", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":            []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties":   []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Context-Properties": []string{`eyJsb2NhdGlvbiI6IlJlZ2lzdGVyIn0=`},
		"X-Debug-Options":      []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":     []string{`?0`},
		"Authorization":        []string{`undefined`},
		"User-Agent":           []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":     []string{`fr`},
		"Sec-Ch-Ua-Platform":   []string{`"macOS"`},
		"Accept":               []string{`*/*`},
		"Sec-Fetch-Site":       []string{`same-origin`},
		"Sec-Fetch-Mode":       []string{`cors`},
		"Sec-Fetch-Dest":       []string{`empty`},
		"Referer":              []string{`https://discord.com/register`},
		"Accept-Encoding":      []string{`gzip, deflate, br`},
		"Accept-Language":      []string{`fr-FR,fr;q=0.9`},
		"Cookie":               []string{`__dcfduid=659b7120785211ecaed6a5158183df4c; __sdcfduid=659b7121785211ecaed6a5158183df4c37874ed17fa4c1a0d743a29588308044f9229446a217f31c602a11fa49438ee4`},
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

//autorequests_whTHctcuAxhxKQF make a GET request to https://discord.com/api/v9/auth/location-metadata
//with 21 Headers and 2 Cookies
func autorequests_whTHctcuAxhxKQF() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/auth/location-metadata", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Fingerprint":      []string{`932960790959493181.999uKkMHr9uFYqRI9m_2Y8vPZIk`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`undefined`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
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

//autorequests_DaFpLSjFbcXoEFf make a POST request to https://discord.com/api/v9/auth/register
//with 24 Headers and 2 Cookies
func autorequests_DaFpLSjFbcXoEFf() {
	data := `{"fingerprint":"932960790959493181.999uKkMHr9uFYqRI9m_2Y8vPZIk","email":"harparser@gmail.com","username":"yothatsmeelol","password":"heyhowareyou667","invite":null,"consent":true,"date_of_birth":"1990-12-05","gift_code_sku_id":null,"captcha_key":null}`

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

//autorequests_RsWxPLDnJObCsNV make a GET request to https://newassets.hcaptcha.com/captcha/v1/76b5200/static/i18n/fr.json
//with 16 Headers and 0 Cookies
func autorequests_RsWxPLDnJObCsNV() {
	data := ``

	req, err := http.NewRequest("GET", "https://newassets.hcaptcha.com/captcha/v1/76b5200/static/i18n/fr.json", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://discord.com`},
		"Sec-Fetch-Site":     []string{`cross-site`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
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

//autorequests_lgTeMaPEZQleQYh make a GET request to https://hcaptcha.com/checksiteconfig?v=76b5200&host=discord.com&sitekey=f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34&sc=1&swa=1
//with 18 Headers and 0 Cookies
func autorequests_lgTeMaPEZQleQYh() {
	data := ``

	req, err := http.NewRequest("GET", "https://hcaptcha.com/checksiteconfig?v=76b5200&host=discord.com&sitekey=f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34&sc=1&swa=1", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"Cache-Control":      []string{`no-cache`},
		"Content-Type":       []string{`application/json; charset=utf-8`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://newassets.hcaptcha.com`},
		"Sec-Fetch-Site":     []string{`same-site`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://newassets.hcaptcha.com/`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
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

//autorequests_YzRyWJjPjzpfRFE make a POST request to https://hcaptcha.com/getcaptcha?s=f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34
//with 18 Headers and 0 Cookies
func autorequests_YzRyWJjPjzpfRFE() {
	form := url.Values{}
	form.Add("hl", `fr`)
	form.Add("motionData", `{"st":1642505628972,"mm":[[50,39,1642505629050],[24,47,1642505629067],[8,55,1642505629100],[21,51,1642505629800],[21,50,1642505630021],[21,47,1642505630044],[22,44,1642505630060],[23,42,1642505630076],[23,40,1642505630092],[24,39,1642505630108],[24,38,1642505630124],[24,37,1642505630141],[24,37,1642505630205]],"mm-mp":46.2,"md":[[24,37,1642505630391]],"md-mp":0,"mu":[[24,37,1642505630396]],"mu-mp":0,"v":1,"topLevel":{"st":1642505627521,"sc":{"availWidth":1440,"availHeight":833,"width":1440,"height":900,"colorDepth":24,"pixelDepth":24,"availLeft":0,"availTop":23},"nv":{"vendorSub":"","productSub":"20030107","vendor":"Google Inc.","maxTouchPoints":0,"userActivation":{},"doNotTrack":null,"geolocation":{},"connection":{},"webkitTemporaryStorage":{},"webkitPersistentStorage":{},"hardwareConcurrency":4,"cookieEnabled":true,"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36","platform":"MacIntel","product":"Gecko","userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36","language":"fr-FR","languages":["fr-FR"],"onLine":true,"webdriver":false,"pdfViewerEnabled":true,"scheduling":{},"bluetooth":{},"clipboard":{},"credentials":{},"keyboard":{},"managed":{},"mediaDevices":{},"storage":{},"serviceWorker":{},"wakeLock":{},"deviceMemory":4,"ink":{},"hid":{},"locks":{},"mediaCapabilities":{},"mediaSession":{},"permissions":{},"presentation":{},"serial":{},"virtualKeyboard":{},"usb":{},"xr":{},"userAgentData":{"brands":[{"brand":" Not;A Brand","version":"99"},{"brand":"Google Chrome","version":"97"},{"brand":"Chromium","version":"97"}],"mobile":false},"plugins":["internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer"]},"dr":"","inv":false,"exec":false,"wn":[[801,754,1,1642505627563]],"wn-mp":0,"xy":[[0,0,1,1642505627564]],"xy-mp":0,"mm":[[343,584,1642505627792],[339,579,1642505627817],[336,573,1642505627834],[331,564,1642505627856],[326,554,1642505627897],[317,526,1642505627937],[316,521,1642505627954],[299,478,1642505629037],[248,499,1642505629109],[239,505,1642505629130],[236,508,1642505629146],[234,510,1642505629162],[232,511,1642505629178],[232,513,1642505629194],[231,513,1642505629211],[231,514,1642505629227],[231,516,1642505629250],[231,518,1642505629267],[231,520,1642505629288],[232,525,1642505629306],[234,527,1642505629329],[237,530,1642505629347],[239,530,1642505629371],[242,530,1642505629387],[246,527,1642505629411],[250,522,1642505629427],[254,518,1642505629443]],"mm-mp":28.15254237288137},"session":[],"widgetList":["02ohj07em6ru"],"widgetId":"02ohj07em6ru","href":"https://discord.com/register","prev":{"escaped":false,"passed":false,"expiredChallenge":false,"expiredResponse":false}}`)
	form.Add("n", `cbd88b740c3569e8388d7168bb76a099c514b408d9747b005e72dd7863fdb17e015d001170d8a179e178051fccb0f00486f812a565ec4b4fa774d54d4f27ec2dc046582f03bc537714ee0c5e21b6abefa44cf74ff3e02642be77a13b685883497c956c99d04472b8f81242ee384243112b0dcf3a4cb30f21a4883208deff7b44b32a60f7d293ff246ae1fc6d8f3c62df420aa0f97b51d7b165caec29354abe465d546aae8ce31a958dfb8c1b5fdd515f9dbee1b49e0e16a5d2b02ef89d4b392e4fce566ab32ecd9e7f304a640cdcb62005551c19bf80ca505d25fbcb6997b14634438df4bd04f5501e7d273bd7721e94a24cdb6e7e1b1a2aecf81fc25e9eb3a84fa9d2d07908e37a2356ca46d4b9d1c3902ff0cc049c6c88d089e517bbb5bef3099d5e7f17c2b58822d773013efb778a1c9eea6af199193be88a3384725d98eee5d504b9c8e03f320d16af74f083b662b96503984580c0b1a93abd862c177db82f8465b897a58f00b3fb258c0e51847efe91645b39f1b516f79b37a40e895e80fd8ee487ddf5d4cfccb78e37097192a1d3affe1b7066954d0ed5cbe154de8996a60b4ace5b6349087018b81073d879566bc62669d369734e96367b690e6aba493eff1ebc366b86e756b861a3b99a5205aeba745450a7a2d493eaa6d6881f81d49d051cff1271ecb6970125b6e3d8092782b7f9859ba37093e08a8d2f4b7e54794c5022a06d068764799a7d3ee2fc94023800863a3ced4b04f4816a46c8aa5ba0d0769f548490fb6517e162f71bf025084da09b22a24a676e59347fe69912a86076ae3bd081f395244481856f8cc0bedaf065ba2fc556d3a51f2d8bb0dfc5db6b44d9be411738cc791544405c5f184d7078122bc22c274bdbcd37878b59d1976974937a58d81391aad2802162c2a3f649f5539046508b5ca3c4df4aeece84b33374a6ed577f86bb3712eb2be672aa8fa7148f4d0ff3eeb2de3b7e4562c9697000d1f5ebe9149bc8c9bac3512637dfa1ea61efdbc0a2e8ada8025ac5ec6e2fe96cb38b6844caf8d3616158cfa2a88efbf024bf1f7884490208ea5162a505c7959981c6ce85723263ef54c386393c239d9c9d308ae98a7d362ee151742f60cd03fbea300409f79eacaa000c0293e21f44e275cfd4af1c9363235a76a75806b1e42871498587dea39308e3cc2308888880e3cd878b288b3aaf269131f28941b7514b82018b61dfff05ac6da4cff102b65db0720183db5f2adf0f49e28581a13ea156decddf43339554b6d5779e1886f82bb774ada504d4f2dec8b6f2ee91a14d58f34183a258c54b8d93bee8e27942485549012c01c6bda73e7ebd8c744890775732ef0f9ff8081911824e290375ad1fa5d3b7d721b61c3afce10ef0721ee48295dfec762ab5f794e39d5fdc699521b73ad9ffd3a87edeb84e5a73a7c6b7b7cebb7919e2731005ee240ea3f0de6b3c2add767fa8b267a3a62f71aeb9acaa1e1535f2b9d791fdd6482f61d75e9bfecdd154d092c7c54995b50c7639b7024e09935b3c08d3b53b71b66308f86f306451fbf4932fa30049bf1f388783bb94a9534758e1cda58576892daca6ecff46258b57988211dd44bcc6c4e0d3f38d452358cf875ab28b7e9b3e7550b51ed36dc3484d0943790ecc3f6abb36c5e73a6a9cc46e9a95f6e1ccdceef46292646fbc2319a90a9ae22c9faa89abd98b750c0e94ad8f831859997cab40d96822ae01b2a55c181fb97c4e91234729a238941d8eeba3cb7dc3187ba3a35ae4ea9ee1171b3fcd72fd27b192a7959441360cf998fa3f571dec569d039f2dfe603abaaa36a765c432cfd17471e2fddd51acb126807b024fdaa553c1c0f923485a5a6ea205535d267ff82923131de332543a778ad4705448426b3eef7737bf61f97c3b0ca106bdd32c913a9e301c05953737c65b1dd950cd53654762754fef704f8d43e33d9009a1990f8ed9ff1baae9c87e8a0fbda47b25a27d2877e5d9296e8e917d603a68c531f63f9130193abaa35eb840557b3a1f5617bb2825d71286133a93e7971bbc5549d5ccb31eb2d5cb9eedeb27efeb7eb8a803e367ad8a508e14d7b6344e7b32218d66c4336ce0d2d7a486410e2082f576398ad84f1ac6457efbc3c671077daa5343a81d86d17ee924e66c61dd20169158f25632f3faa85fbc00ed6066e020060e442e0cda8f52ecdd275775f177cad71f773568ab2744dc42f258b6db819ad79075ced972ca0d692ef3e4d6a468cf4bc6f87f7d45d271be4a9078f837da47cb2f94ceb6beab225b1387db06f50de7bb8a3885f322c609d91e6306a854e2befbfcc97ec288e99e5848b29a2b7811bf8d94a483f91df0a898d56f3104776c4d14b679b6b30dda6314d506f8fb30e3c2c52e0c7ed9a7910996c1b4cb3ff24fa658d79608e0078ea45c7d85c9a74da5ccb2b01235f3002e24438e954d74b15da7f2b53c85d55bfdb5cd83552bafae50cf1ec270c765dfdbb20b9fa30649bfa6714e93c125ef49f93807b0f598dc4876e0a0ef764f47447ab90203a36ba341b672e654f0eeac4048c2770c42b5e290cf84990f2424d20291493d26836fbd55ca0711d521e6c17e46a3e141725c00`)
	form.Add("c", `{"type":"hsw","req":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzIjoyLCJ0IjoidyIsImQiOiIrd2N3Mi9xTXJLZ3JLbCtKTFhtb0w3VVpzbldxZHFOQ0hpbmhETWNZUmZxbWZiL01mWTVwdXBqNWcxTWdLTTJDZ1QxT0RXS1BlYUR2Um9haFUvdnJ2S2dkbGNkVGFTMDlRdE9KV2FVMGovaXRXU2NaTzlEbEQ0N1F4ZkRvRnBYMmZMYkZjOGVSSERlRWpuVXZMa1JwV2FiN0lXUm1CTlIyczhhbGNreW42K2lkbUxycHhOZllXYXAyM1E9PTgyQzR1NnhZUG4zdkc2VzYiLCJsIjoiaHR0cHM6Ly9uZXdhc3NldHMuaGNhcHRjaGEuY29tL2MvNWQ1NDAzNTUiLCJlIjoxNjQyNTA1ODY4fQ.9TWF7BFtRELSyKUZBF8eDSW1v6httGa5aTAl9teSPCE"}`)
	form.Add("v", `76b5200`)
	form.Add("sitekey", `f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34`)
	form.Add("host", `discord.com`)
	data := form.Encode()

	req, err := http.NewRequest("POST", "https://hcaptcha.com/getcaptcha?s=f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Content-Length":     []string{fmt.Sprint(len(data))},
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"Accept":             []string{`application/json`},
		"Content-Type":       []string{`application/x-www-form-urlencoded`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Origin":             []string{`https://newassets.hcaptcha.com`},
		"Sec-Fetch-Site":     []string{`same-site`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://newassets.hcaptcha.com/`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
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

//autorequests_gmotaFetHsbZRjx make a POST request to https://discord.com/api/v9/auth/register
//with 24 Headers and 2 Cookies
func autorequests_gmotaFetHsbZRjx() {
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

//autorequests_AwnwekrBEmfdzdc make a GET request to https://cdn.discordapp.com/bad-domains/hashes.json
//with 16 Headers and 0 Cookies
func autorequests_AwnwekrBEmfdzdc() {
	data := ``

	req, err := http.NewRequest("GET", "https://cdn.discordapp.com/bad-domains/hashes.json", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://discord.com`},
		"Sec-Fetch-Site":     []string{`cross-site`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
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

//autorequests_EkXBAkjQZLCtTMt make a GET request to https://discord.com/api/v9/users/@me/billing/user-trial-offer
//with 20 Headers and 2 Cookies
func autorequests_EkXBAkjQZLCtTMt() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/billing/user-trial-offer", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_TCoaNatyyiNKARe make a GET request to https://discord.com/api/v9/users/@me/affinities/users
//with 20 Headers and 2 Cookies
func autorequests_TCoaNatyyiNKARe() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/affinities/users", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_KJyiXJrscctNswY make a GET request to https://discord.com/api/v9/users/@me/survey
//with 20 Headers and 2 Cookies
func autorequests_KJyiXJrscctNswY() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/survey", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_NsGRussVmaozFZB make a GET request to https://status.discord.com/api/v2/scheduled-maintenances/upcoming.json
//with 16 Headers and 0 Cookies
func autorequests_NsGRussVmaozFZB() {
	data := ``

	req, err := http.NewRequest("GET", "https://status.discord.com/api/v2/scheduled-maintenances/upcoming.json", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://discord.com`},
		"Sec-Fetch-Site":     []string{`same-site`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/`},
		"Accept-Encoding":    []string{`gzip, deflate, br`},
		"Accept-Language":    []string{`fr-FR,fr;q=0.9`},
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

//autorequests_sbOJiFQGZsnwTKS make a GET request to https://discord.com/api/v9/users/@me/affinities/guilds
//with 20 Headers and 2 Cookies
func autorequests_sbOJiFQGZsnwTKS() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/affinities/guilds", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_mVoiGLOpbUOpEdK make a GET request to https://discord.com/api/v9/users/@me/library
//with 20 Headers and 2 Cookies
func autorequests_mVoiGLOpbUOpEdK() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/library", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_updOMeRVjaRzLNT make a GET request to https://discord.com/api/v9/applications/detectable
//with 22 Headers and 2 Cookies
func autorequests_updOMeRVjaRzLNT() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/applications/detectable", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Cache-Control":      []string{`max-age=0`},
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"If-None-Match":      []string{``},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_XYeUCWKsXbGyRAO make a GET request to https://discord.com/api/v9/users/@me/billing/country-code
//with 20 Headers and 2 Cookies
func autorequests_XYeUCWKsXbGyRAO() {
	data := ``

	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/billing/country-code", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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

//autorequests_mBTvKSJfjzaLbtZ make a PATCH request to https://discord.com/api/v9/users/@me/settings
//with 23 Headers and 2 Cookies
func autorequests_mBTvKSJfjzaLbtZ() {
	data := `{"timezone_offset":-60}`

	req, err := http.NewRequest("PATCH", "https://discord.com/api/v9/users/@me/settings", strings.NewReader(data))
	if err != nil {
		// handle error ...
	}

	req.Header = http.Header{
		"Content-Length":     []string{fmt.Sprint(len(data))},
		"Sec-Ch-Ua":          []string{`" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`},
		"X-Super-Properties": []string{`eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJmci1GUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzEyXzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85Ny4wLjQ2OTIuNzEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk3LjAuNDY5Mi43MSIsIm9zX3ZlcnNpb24iOiIxMC4xMi42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjExMDQ1MSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},
		"X-Debug-Options":    []string{`bugReporterEnabled`},
		"Sec-Ch-Ua-Mobile":   []string{`?0`},
		"Authorization":      []string{`OTMyOTYwNzkwOTU5NDkzMTgx.YealoQ.UWtSDLERNzzkos0ietKVWGtd4Gc`},
		"Content-Type":       []string{`application/json`},
		"User-Agent":         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`},
		"X-Discord-Locale":   []string{`fr`},
		"Sec-Ch-Ua-Platform": []string{`"macOS"`},
		"Accept":             []string{`*/*`},
		"Origin":             []string{`https://discord.com`},
		"Sec-Fetch-Site":     []string{`same-origin`},
		"Sec-Fetch-Mode":     []string{`cors`},
		"Sec-Fetch-Dest":     []string{`empty`},
		"Referer":            []string{`https://discord.com/channels/@me`},
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
