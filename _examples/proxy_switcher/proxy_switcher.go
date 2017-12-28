package main

import (
	"bytes"
	"log"

	"github.com/AlexanderChen1989/colly"
	"github.com/AlexanderChen1989/colly/proxy"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Fetch httpbin.org/ip five times
	c.AllowURLRevisit = true
	for i := 0; i < 5; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}