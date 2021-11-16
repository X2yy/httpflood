package main

type Ataque struct {
	Url string
	Host string
	AtaqueMetodo string
	PostData interface{}
	RequestPorIp int
	Cookie interface{}
}

type System struct {
	Banner string
	HTTP2Timeout int
	Ataque *Ataque
}
