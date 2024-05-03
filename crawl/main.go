package main

import (
	"crypto/tls"
	"fmt"

	// "net"
	"time"

	// "log"
	"net/http"

	"github.com/gocolly/colly"
)

func tls_transport() *http.Transport {
	result := &http.Transport{
		TLSHandshakeTimeout: 30 * time.Second,
		DisableKeepAlives:   false,

		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_AES_128_GCM_SHA256,
				tls.VersionTLS13,
				tls.VersionTLS10,
			},
		},
	}
	return result
}

type noteLayer struct {
	topNotes    []string
	middleNotes []string
	baseNotes   []string
}

type metaData struct {
	name       string
	brand      string
	accordList []string
	pros       []string
	cons       []string
	notes      noteLayer
}

func main() {
	rawURL := "https://www.fragrantica.com/perfume/Tom-Ford/Tobacco-Vanille-1825.html"

	c := colly.NewCollector()
	c.WithTransport(tls_transport())

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Safari/605.1.15")
		r.Headers.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.3")
		r.Headers.Set("Accept-Encoding", "none")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.8")
		r.Headers.Set("Connection", "keep-alive")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL: ", r.Request.URL, "\nError: ", err)
	})

	var m metaData
	// // Set up callbacks to handle scraping events
	c.OnHTML("div[id='main-content']", func(e *colly.HTMLElement) {

		fmt.Printf("NAME\n")
		e.ForEach("h1[itemprop='name'].text-center.medium-text-left", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Text)
			m.name = el.Text
		})

		fmt.Printf("ACCORD\n")
		e.ForEach("div.accord-bar", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Text)
			m.accordList = append(m.accordList, el.Text)
		})

		fmt.Printf("PROS\n")
		e.ForEach("img[src='/ndimg/Pros_icon.svg']", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Text)
			m.pros = append(m.pros, el.Text)
		})

		fmt.Printf("NOTE\n")
		e.ForEach("div[id='pyramid']", func(_ int, el *colly.HTMLElement) {
			var bucket *[]string
			el.ForEach("/div[1]", func(_ int, subEl *colly.HTMLElement) {
				fmt.Println(subEl.Text)
				switch {
				case subEl.Text == "Top Notes":
					bucket = &m.notes.topNotes
				case subEl.Text == "Middle Notes":
					bucket = &m.notes.middleNotes
				case subEl.Text == "Base Notes":
					bucket = &m.notes.baseNotes
				case bucket != nil:
					*bucket = append(*bucket, subEl.Text)
				}
			})
			fmt.Println(m.notes.topNotes)
			fmt.Println(m.notes.middleNotes)
			fmt.Println(m.notes.baseNotes)
		})
	})

	// Visit the URL and start scraping
	err := c.Visit(rawURL)
	if err != nil {
		fmt.Print(err)
	}
}

// 이름
//<h1 itemprop="name" class="text-center medium-text-left">Havana Aramis <small style="white-space: nowrap;">for men</small></h1>

// 노트
//<div class="accord-bar" style="color: rgb(255, 255, 255); background: rgb(119, 68, 20); opacity: 1; width: 100%;">woody</div>

// 장점
/*
 <div class="cell small-12 medium-6" style="border: 1px solid rgb(207, 249, 207); border-radius: 35px 0px; box-shadow: rgba(216, 216, 216, 0.73) -11px 6px 13px -2px; font-size: smaller;"><div style="display: flex; justify-content: center; margin-top: 0.4rem; gap: 5px;"><img src="/ndimg/Pros_icon.svg" alt="Pros" style="width: 1.5rem; height: 1.5rem;">
 <h4 class="header" style="background: linear-gradient(to right, rgb(108, 143, 109), rgb(207, 249, 207)) text; color: rgb(131, 166, 196);">
*/

//노트
//<div id="pyramid" class="grid-x grid-padding-y"><div class="cell">
