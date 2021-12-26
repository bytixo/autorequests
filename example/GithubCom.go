
func GithubCom_XVlBzgbaiC() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/users/bytixo/hovercard?subject=repository%3A441991500&current_path=%2Fbytixo%2FHAR-Parser", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Accept":	[]string{`*/*`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=K8OKDrJeRaD9Zm4PJdEknnDI5DU24zHj1D38oBdGIJNbsUPE%2F6K72lj9GiekH1GsovVoJCxO%2BkSLOXxWBP08mNyt5po8ZrNBk3M%2Bkp0yj9LZhMFa3UZcoaO416d%2Bm3d%2F2yPOpTRbj%2BOGtaGpYior%2Fx42pX0HZmIcXfb%2FAJRCb%2B4W0OP%2Fl4ZJfBYFieYTwpGuZqvD4U42WiNpv3Bq4cX%2FwbHInin2QbS1RCiQbwriqOoinNLgQRIlavl0ifqp%2BJDThum2Cyu2vCqT1GCZo1J8OPgPniZhaUlqBh6V%2FEQM2CMAtYbnTBJE1TjLW47nP%2B3qjnbUQM453Mw2kSxdWNdl7eTda0dMbVjQbeCCBw23hOyBOMzQOQk0uikt2acKYvxsW0GEM506cZjqTUrYZOn77%2B%2B1%2FbbTQw7QyQJaHh7%2FO5N606PqJkaiDDeTflgMXHr3ZpcAzZuP52fUeO96eoKclvmISd8ooxfYRY3c2BVDRqXBYs%2FKlWABWwL%2FVP9%2FMUp%2F5lWIwSbkWMy59S5dDjUttcmBsK2E15BUEDrIFVM5SGXW8vpMTRYpribs7S6ZxXaRbmfuIXpRM3mVT0nlxdV47z6VnZYI4D7vbwTxax8ggA64PczdF3LDMnuhbzLiJcQXZxqV1vr9Lj%2FCzaRdRLZmMd6y70JI%2BlvoOCxYwL%2BV%2F9kRMhplTgcG93B7Q9NdLDveBbNcCfLhokosxVfck9olZO8zsGS2%2FEcdcr7%2BLLIAmaY6w3HzVT7hTbNqNL25LGNkpydTX5YmH9WrrEOO8mvtzm%2FJuRpZcH%2B71T7MAXyPmhCHgmhnL%2FEGdWg16eptyhXlhAMwbRJlvPjOIJxo6QG1lfjdpwt8usW9po31Zsa9jzfPqcQuEVa762rvzhlM7uPgoM37KzM9Sq1DHFM9KKG6LpkRUe0Iam5%2FDM%2Bjmp1PA%2Fg%3D--xwIAf2ZNH8qi1MZX--e6iAOHbMdQrK9DEeuXx%2FEw%3D%3D`},
		"If-None-Match":	[]string{`W/"2dd3561fd0ce6221e8578f35d166095f"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_MRAjWwhTHc() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/blob/main/main.go", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"X-Pjax-Csp-Version":	[]string{`9ea82e8060ac9d44365bfa193918b70ed58abd9413362ba412abb161b3a8d1b6`},
		"X-Pjax-Js-Version":	[]string{`996e21374117f01b7df2c1f75047cbdb09449e0572a60b7722454a61ba4735f6`},
		"X-Pjax-Version":	[]string{`2a0ed398d891434ebc5cd37a121dd8663578609e0ecf0d1b542fb263bb3ef861`},
		"X-Pjax-Css-Version":	[]string{`8c06702a824b73d05c6dcbe160dc26951e0c3af9f2b6bd2c34681db7ef946a74`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Accept":	[]string{`text/html`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"X-Pjax":	[]string{`true`},
		"X-Pjax-Container":	[]string{`#repo-content-pjax-container`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser/blob/main/main.go`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=FrZzGOdn7Os5QCu1Avxh1VQGCnW8gBwhIWvDCnA%2FsU8KJQv54LI3Cg1Vj9B8dheh%2FEHVl9GQfY%2F96NaC%2BCSKMFDuBtLMhGj67RxAC7Srn90njz4AVVDWnkGDfJSzcuUvsQe346i5kJZHgFCkZnoV7rNQAJwxGl6%2FtTD9poTfecrMzSplNWo3ljv9K6mEK4ZWG5v753fb6P7PzIzIKKMhEl8JIpgZxrHDRbqc8fm2LU2hPSyTpxw2BnjZPEjMJ3YFsr3Bz8Zcmwsas%2FRJRkLH7eAPqbih9m4LNLmaqGNJd7wjb%2BNQ6ZmKDM1e141sICyt6TRIEToEi5Lan3lBAS7ULd%2BL%2F0Ii5w3TUNcYadHuFDFwLNZjM5stbYScTkwXAAnpg9WGc7w6A9mESWgS9hSduhwFYE24FITJPOtPvAmBcQ8ADHNDhw697EJA7URNxPtKlFC63BvKu%2BgXED%2BUFbeSf%2FBZwO85jC3tkMruQ1za4nwtorKl1Pe1pS4VbklcskuMMPlRBVCO8y1tUF6hmS2eLlF5nApfaDB4oogCJl%2BauBIoZD9ey3aXL7n8KejPe7i5uNDuD2DRNjPe5w6sokK2%2FgZxXbLhXtKgdso9wOeduUtfpf2oXY2VM2pSA87%2FWvFIBHEW8xyHo4ZS2hv93fSO50Ba6VNl38Mkus3Ahj1EO8XAz9NVzMFt%2FN%2BuZepfrhwmS7vbZ1qO43yUQs%2B5k%2FJdJ1%2FLvmiJzysnS4qrHZ06dOW8dtB7PaAiCU2ktmF58jZv3rQPa%2BdK4Vr7MjF7hSjfy7OCBN4nR%2FFe4yGyfwrWbl4EAbxqRgaIHSOVfUe%2BCHp6QeBpBkqKpSzpc4%2BecQA6BsfCMhWghIUxgQ82avsgXx68ksQmTRfRlx6iiehD0R7zcEZfzXFVuq7O5tnDBA0STfANfH2WKi%2BIGsB2XVq483M%3D--owuIV5zzft%2Ff2wNg--7lN5Fnoxq2e9RDUEv3A7Xg%3D%3D`},
		"If-None-Match":	[]string{`W/"950c53610819084d0b0602887f4a00ad"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_tcuAxhxKQF() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/contributors/main/main.go", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser/blob/main/main.go`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=NdETdYdPOcYM4R01T59fv3EYC2kvuFKsckvnsyTZWmFEUukLw8fbg1eWBVr3Wg0vPIeqo2I9rQBw4lNfKMrgjT8EPbhIfZMxrXtDK%2Baj0ev477TuLUvfbJZRj9KRzieLcyGhh53lUg67oruuZfZXtbu0uvcFTxEwHJA9i03mWbeEPmJzh178Y9H6hiYSxdZnZdZOsddnadGxxb5cqyajsLOn86t6ayXJw3UO55lsVtcdh%2FJCoT%2BpLllNqun2QsqYtvzJVq8P8w6CmzKTHzFt7jVUw%2Fm7yTwiGeCT0zuw36ZGup2PsaixnI7xE%2FJCyD9g67f0HE59YTmx30xdC9SjzveTqkbf7LBWTkOD%2BtTTKNq7ovGtPUZcYdr3E2tgYhPSni5XYP7PeAP28wrFx6%2FbLc%2B0KiaCGoRb9PK5fn6pq6V6dLzYi7eKffjiK%2FFfeect2qIElgefHqw7mxAa%2FTA%2B67jphHgFfNWKcecVMnQLB4Jp3ctQ5er7AOnRLT2jq%2FO0wIp4sK7U2%2FFaAwRGlyd5EeCyx8UEwc094MarIBUi0vujAL9rasE5yYkl7bhc7YL2xICEqIQ4ISQZvzbRWva6Aoo7DbF17ho88jsqfFLKHZAOGctN%2BhDkOV6ofV9sWl%2Bn%2Fpjsg2jHpe7dvoWtNtFh4q9oyhrBVj5HKv%2FwHVcIxXCbWdzel%2FUVAeV4xYlGBedK8hmFMQscsYKo3D6muowcXa%2BYT88PMPGHcwhebPiadYnMmLgzmLt8bf94p6ajtuLkXzK8QktntjGwmcxFbpx8OChmPwa5A78u7%2F54r76fkKW8E8DOHxL21Z0LMunkCnuVfbTYmLZ8V8RVCPfVYICex8U8n6JoETa6TH%2FUwXnlW1skl7uZUq5bMghmfjKrakh7VvgF5HDhZ2lJgEAPu8zqpfRHjWo8BWd2FJe3UIpTAlw%3D--5XjcwdrIaQvTl74c--0H8R5%2BEOfmarx7jaBSZ6Lw%3D%3D`},
		"If-None-Match":	[]string{`W/"ebb9e319892faf647e2c925dddfeaaaa"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_DaFpLSjFbc() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/commit/2f44a6d5263e9b31845e1f2a9892f2619d28845a/rollup?direction=e", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/fragment+html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser/blob/main/main.go`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=8%2BTCQ6iPYzjiyXrhuOAMrh85Uv76UlltI8YOcUJ%2BxTdi%2FFDeTUocsPHt9Lk9HafdNfLGVsX1slvnHn3wbnnxt8uuuM%2FqbeTb19jQbqQX0i1Zzy1wPZTkxFmoCK2UuOmzY900A9tjRhZYa04uWcRbczidf1FKJW1yZ64X7VZod0z6TRWOsBfgpWMxm9ZiCtX9QG8pC8iHzgAxjVRY4qLlcxznwsjYgJx1uEhGxZ38AHscLpeBIgLf2o7SY08SLWthcWzQGNUKxIvT9ICcNuVMkjH7QpxiNoYMzv4xhViwYqZh0J2z9kLNrSWA%2BgOTJELQjQBi15wkMPpm6ssDjUTT5TyjKn5eBDWNKI3cI0i0QaNWbLzqha5o38GSqYi8qPHKcehNToQWzqr1alUVdU%2B2uaob6KFUn%2FGGNxcRnwhMz673QjPcAR%2FpVzuaEzu2xjpC9fbIdCJ%2BSSHShEnlYytRtgQVrViCHcSakVYqNFfG1vDpGQK5eS5dgQyJeQunYU0TK9Cw0LOhiJGC3pmvYp0CtY63jPxzqMq%2BWXWX774ciQ3dna%2BpPHUmg8XcsDZzMo8OdFjptEv0x6gHsyHrI7Ul5OMAwbjYZWSqAUpLLT1v3pgcjgldS4LhkUIQEVPee5fXxqQaifa4TNOc37t285HREizwoJamx3jZfDomOHoAapnH75t1lYgj5HpfBRYqCqkOUvYiP0ivGo%2BaT%2BxTwvite52VUDhua%2BbCcQUbM2wnx1P%2FAKVqKk6lQ5zkyaF3KBoCaJnSasuI%2FuI8vfgQj72tV4U9j9aZ2jICx89Mzv5fjrADgaLXoLux%2FiKgMQidZWrnBjmJd3k4EOeqStBi%2B1UWNdeunxm0mFogYQil9z6X9KzDUuuwDbjzMHskz4JwBd%2BLMCsqU2b8ZIXNqnH2ChbCnxiEyP4CUGp2xlPmCktqoS0%3D--o1J8GVwmJu6kSeRB--IucKhMCNk%2BLLdeAa8%2Fsa0Q%3D%3D`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_XoEFfRsWxP() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/find-definition?q=main&blob_path=main.go&ref=main&language=Go&row=0&col=8", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Accept":	[]string{`*/*`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser/blob/main/main.go`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=O0wP048VDg9lncMC4eJRB%2BFmm5I4zu0KtsIKw0g%2Bq8%2FFtg09wu105kfFFEetGwgUbWpFiShgF%2BoBslNTri7GsM72ez7nDG0StYkh54se0NkVyvoh9aWeS8Tr9Ww4TNc%2BOdrbpQ91%2BsCYV6awjCNSiKplOK1VhWHAC0NznbHqRo1Ima%2ByFYJQ524VJTvx3NI76FUo%2FkO4TrasnOGwM2B9SxiNA3bkKcnHy0TQp3VqVjGWz0F1Mi1Z6FQnxx4UZQ8ya7Qhj0bigM4Zii9QW5sM96xoQS3%2B%2FIx526n4ICe8RyazZwUfq2CQr5JhjwThC1ufYS6X1SzgrJtM2%2FTorndCEn7M8jqBXLZPE5BY72SXTajuWmsIRwYhYegaF%2BxEW2fZthKsJLA9e4FnSpimcTVx5b2aZoVlhkRTUO6PmD56bRbcf3fg%2FGEEIavlx%2FugWqHasMhhYcuhVSwpmRkO8ZlN%2B2uHR%2B4hU9VVVLa991qGB6xm65ml2DXu0gZeU2x7Z1YQ6Qtojygxco5Gj7tjZJ6z13uGhRUIRyfUV0T6JnJu1jCiDLLo95%2Fn%2Fu1iUlR551cm2Geq%2Fy2CT3Dycqh4Amtm%2BsPIOG6OKVRW089kPqIzVE6P0CAgpPjfTW4CDdUuO%2FTzENhwLasYcW1nz3IP%2B3jFUoecUcgujYahSPPv4BlNbeEMMW9P3H9lAP0UALoc24hi6aOmMILz%2BIUh8d4bHMpT5XlLZ9bg4f6IluOYtR%2BEtITMpfHaETAghxHHb2iznM5xjSw5C%2FimoI87SzREXRGaF0qH5enQOfvEHt%2BDBV31tIMRlWvtNBTOcWF7BxdsxnfhufR4nr2vwwC13Rz%2Bo6s9B4Rq%2FFJDt1iMWC%2FXjc%2Fbc2vq6LAFTqB7uVQhVYgR5DgJXhDpZHkI8cbVSVj2TExgN2wFgvbUycpD9IcdCvLY6IM%3D--s42V%2B3PLbaUJHadF--F7IfgvlrNe8KXYoGrUHC9A%3D%3D`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_LDnJObCsNV() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"X-Pjax-Csp-Version":	[]string{`9ea82e8060ac9d44365bfa193918b70ed58abd9413362ba412abb161b3a8d1b6`},
		"X-Pjax-Js-Version":	[]string{`996e21374117f01b7df2c1f75047cbdb09449e0572a60b7722454a61ba4735f6`},
		"X-Pjax-Version":	[]string{`2a0ed398d891434ebc5cd37a121dd8663578609e0ecf0d1b542fb263bb3ef861`},
		"X-Pjax-Css-Version":	[]string{`8c06702a824b73d05c6dcbe160dc26951e0c3af9f2b6bd2c34681db7ef946a74`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Accept":	[]string{`text/html`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"X-Pjax":	[]string{`true`},
		"X-Pjax-Container":	[]string{`#repo-content-pjax-container`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=XrCcxSXvROGEmGaKn6cunDA53JqUWwpfnBidgNUtVYb3y0aup25EWBDTf78t1IfDVX%2BlFt%2F3gVkSj%2Bd1AeyXQD9uZmXscPdCEFIOwcIUD9L%2FIHXMPdMXyZSY8fjBmy%2FiY6dlmbNrsGdmjjOBIalbB%2FTp%2By9DreZNSJUVhEOWn%2FcPICVwzqpd8ZrvFlhDmirO2oBQGJ%2Bu7HngBZMzCRmzkQotrU5UuZVd1sn0xZ%2FPEZinr4UVBydCySkgRhDRoGJFGnFL64sRvSqeTTLDAU6tSWGhqWtiBx%2B8cAgeFdUnQDHDKhvPvDmCUjmF2zXPGw5A63O8xnnYjVj2ZHi3nX%2BIWmti%2FYlJsTUZsX0ywiwKpmeWFtUbelYPGxkTgzlLGlRqWCAHPFnkePAyKawTytMIjeig66BLh6%2BPQFLhRBeFPnQgYsqDfwAimC2mppHPXIHsbD8bemZtS%2F5ebzhg70%2F6VUD8w9bH%2BvXn2ErQTJibM3zuRCo8zx8Ks95DK6m95zvjZ%2BRSVnt3iWkr%2FVebBqGZYmu%2BsauECj%2BLyqIW79fvkTGxUZiwLl3RLIc8G9iP4wU1D8yLH8YfXe6sRp3qTYImG9Pf0wdkVW9Ml93TaB%2BELLN%2FWggoO9VrxWkdf%2FJNpESNYA9CR24HrURXCQkhkoHQg9CtZItP%2Bj2s6BixVipb4w6VVURpBI8cfortHAwXaotq6Cf0hUbVnJUr26hUTUIlov2j%2BOe3YqZBuIeW7RCIWlwMLwFDOcqNylboqz2YQ1EHc1BnAIyifskjACcI3WoC%2BXVbJWp77mruEO%2F%2FOK4Ih5MmA2peSalCyRc941K7mGCIAwCOIVmmBD7PSl2DrQHvmsSPSULyUrYbd811e8gbeAN1luv%2FZVfeT9%2FT5hZ4fzSJxqG2PYCJQQsMLFEjB1TKshoHSCVHcOopeBQOXbdNbsE%3D--6MjkUJwpPHQS3kVu--%2B9Q6n21hCFm5FtucgSB9PA%3D%3D`},
		"If-None-Match":	[]string{`W/"0917918798dfd3a92add86d6159d1e36"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_lgTeMaPEZQ() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/show_partial?partial=tree%2Frecently_touched_branches_list", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		"If-None-Match":	[]string{`W/"fb59b572f9aa02d43187ccbf84b1bcda"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_leQYhYzRyW() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/overview_actions/main", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		"If-None-Match":	[]string{`W/"24544f918f107172bdc3fa59a92e4f0a"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_JjPjzpfRFE() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/tree-commit/fad3c7c5c228abd0451b0a0865b4728d19bd9a75", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		"If-None-Match":	[]string{`W/"5a60ad71baaedfb16ed079f3d43927f3"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_gmotaFetHs() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/file-list/main", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		"If-None-Match":	[]string{`W/"f397fe5038d112a4314813c4f5f01406"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_bZRjxAwnwe() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/topic_suggestions?async_topics=false", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_krBEmfdzdc() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/hovercards/citation/sidebar_partial?commit=fad3c7c5c228abd0451b0a0865b4728d19bd9a75", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_EkXBAkjQZL() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/used_by_list", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/fragment+html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=iY%2BCkbRhADoMUtL3cNNk7777enkTZBqVlmnhEpKe1JwP3nK0MK4KskTmijJueEp6briHjir%2BhGLcjvlA7RAHpLQfdSNotsJViWqPsO8QFtffxGJIPsXqSvBTGeuuFRyV%2FcbtCyJiN5pe%2BHRijUse9SafOMghBrrP1tsa2pDGgNE%2Fh6Ex7XpLfANZAQWOdnTBrj9UgobApStoiu%2BAhJercu2zNeexT8dlf0B8kFlN%2Fw0sPDdTlIOShrR5pa1HJpuBwHMntInj2t%2Fsdy4eZ49tLuBo8MKzb%2BwU64d%2BZPUbVbCWj55peXSYefkXRxMIVB3qknj%2FkdRsVet%2Fc7FWi5IT6O1JVgsaoV0wRshm3MshyN00k1oGgzcCciikolu%2FTLZZ6NvSCSvsuZSQgvvMEDGQ0%2FWuJbAmnMl384w5IN1BosqqoTfR8EZvYVUDBeMC77AfweAdMLM1ubD%2BRJT5O0D2MUOWeJ8DS50yAqNhMVhCIoW1okFkmlsCgGEzDVnPm%2FhzfZAbheXW4I2T%2FDS51o6BloQ9xGK9xHAr9k%2BWECx%2B36E5XjcFmlcXodiLdYG4HTTD3aCMkT5IMZtNuLk2UnKnOh2zwZ1JKZFkLJQLkBIp6ARb%2FzAXLG%2F2ZRFr2dxkyK%2FT9FBB7qtq2WDFzWFW9rKH8y%2B%2BSqP2K0AZTkVD0xjmV9ZfDpnB%2FqSSs4hdXUpnHKqonjEZ7AUd8usG2Dzb%2FpA2Gq8GZ9SRGD%2F794E%2Blze8a5QigZxC56XgB8EYY%2BigsJz2dzz5l1eQYa5a8hu%2FW%2F1nt70kla3PttKbhOAx%2FqQG%2F4jcWE%2Bj6zHaa5oITy5DMNVdrr2Q01bXZuTQ%2BxhULYg9vZARgvJ23%2F9K5sAPuYYlFSa6nkcNo0MB%2FoVjPvuXAwni%2FBJT6p%2Bzro4LzGIebvfWwbr59MspByvTOBDE2I6jJOo%3D--oPxcgW4QDqIeQX9z--kj65x8FXCim3R7biB%2FWJCg%3D%3D`},
		"If-None-Match":	[]string{`W/"75a11da44c802486bc6f65640aa48a73"`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		

func GithubCom_CtTMtTCoaN() {

	data := ``

	req, err := http.NewRequest("GET", "https://github.com/bytixo/HAR-Parser/commit/fad3c7c5c228abd0451b0a0865b4728d19bd9a75/rollup?direction=sw", strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Sec-Ch-Ua":	[]string{`"Opera GX";v="81", " Not;A Brand";v="99", "Chromium";v="95"`},
		"Accept":	[]string{`text/fragment+html`},
		"X-Requested-With":	[]string{`XMLHttpRequest`},
		"Sec-Ch-Ua-Mobile":	[]string{`?0`},
		"User-Agent":	[]string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 OPR/81.0.4196.61`},
		"Sec-Ch-Ua-Platform":	[]string{`"Windows"`},
		"Sec-Fetch-Site":	[]string{`same-origin`},
		"Sec-Fetch-Mode":	[]string{`cors`},
		"Sec-Fetch-Dest":	[]string{`empty`},
		"Referer":	[]string{`https://github.com/bytixo/HAR-Parser`},
		"Accept-Encoding":	[]string{`gzip, deflate, br`},
		"Accept-Language":	[]string{`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		"Cookie":	[]string{`_octo=GH1.1.984421530.1619015101; tz=Europe%2FParis; _device_id=7594121fe3aa43a56b01106b647a6eef; tz=Europe%2FParis; user_session=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; __Host-user_session_same_site=qKzarCka4PVdgvWRYyOqceNKhx3SKiotugDslmFtW0cxkVMM; logged_in=yes; dotcom_user=bytixo; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_colorblind%22%2C%22color_mode%22%3A%22dark%22%7D%7D; has_recent_activity=1; _gh_sess=E7oUyybLYUHxWrA7ZPXr3yXGSmZ3zq%2FmvCfNcYnDGhL3WYB%2Fmk2oTrSblCYTPuhUcopVuUUFNDTn1uSWhUMR6X5z4%2BzXVAhds%2BrT4djFPlZlV1l8frDOvi9nnt9wZtS6rbO1weuzOPsqmJ5Zgq5svUmFQrFlIHgRhcEVh4%2BV7il0A70KI45ScDaGSqKjpj1Pb2%2FCWcbBw1a8b3hMqg%2F5fNiVhTbWFXqw72skeKgoAAKaRq5I1gMQEYYGF7PpNwl9w75BoytX344uiM9rD2xIvk88mM6Q87PIcph%2Bx83NJMZymkCmujGHoAPosadE5aA44Z%2BXC7233tVN9WrNiTA3x6yT1N7RVJflzDslY4rP46cEx%2BmD7pMMp0BBTWFuEOMdKr3zovmX0O6C9oEjirt56N%2FC5PNW%2Fp0cKeg%2BzRAV3eu2CEhuIDNb9GM%2FyIrieRWZC8aqyCmycvYvNd1zlfsh3Ik2WvLFgnv7AkJnGS3ApptCDbwz6nN3fp3f0IGxbdhECOOJsVmvOB%2BM6p0vDpXxhNvtX%2FiuELxO8a%2BofgytvZjbpZoysT4vBKoH4AjEZGJYQY%2FkfQWZKA0zfNYcc8jGoQ3qA6u9s3Zk9srZt%2Bp%2FIzDyy%2BpZTh8JhTIJwGS3YEVRewnp1oG8l875%2FKgtUglFsHOJjtICOnBvrHv%2FkqAnPka%2BWbDgXXjw6A0FKbepgqMvbtSap4US4fyJNKUmly1bEtIvVBg5WS%2FWAin8afFAR8Pu%2Btw2TIaKPJwaAh3X8nKeVsGh3ozaQd6nk8%2F2D6iMGVNViOWnYa9mq5aR0iaWV%2BE1WgYWr0833ZiHBBrP8RcOpisWqsEh%2B9uGeaV%2B5R8NDN0dJVwLxn2N0Ki8rueroODFE5AD%2Bj1pL0DCZK0HyN1bB6aaIzJX%2BFL4R%2BT%2BAFBLcpsBTN7zjhbSZzwIPvnHn88%3D--YhFGZDjzMI1OfmQo--RO%2BWwOwi8iBB7v8V%2FhpPeQ%3D%3D`},
		
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
}
		
