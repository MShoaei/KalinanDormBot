package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getTermID(doc *html.Node) string {
	var value string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "option" {
			for _, attr := range n.Parent.Attr {
				if attr.Key == "id" && attr.Val == "TermId" {
					for _, newAttr := range n.Attr {
						if newAttr.Key == "value" {
							value = newAttr.Val
							return
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return value
}

func reserve(cookie, sessionID, applicationCookie string) (success bool) {
	jsonRequest, _ := json.Marshal(defaultResReq)
	fmt.Println(string(jsonRequest))
	req, _ := http.NewRequest("POST", "http://dormitory.sutech.ac.ir/Dorm/SelectRoom", strings.NewReader(string(jsonRequest)))
	req.Header.Add("Host", "dormitory.sutech.ac.ir")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:66.0) Gecko/20100101 Firefox/66.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Referer", "http://dormitory.sutech.ac.ir/Dorm/RoomReserve")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "close")
	req.Header.Add("Cookie", fmt.Sprintf(`__RequestVerificationToken=%s; ASP.NET_SessionId=%s; .AspNet.ApplicationCookie=%s`, cookie, sessionID, applicationCookie))
	req.Header.Add("DNT", "1")
	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	success = !bytes.Contains(body, []byte("ArgumentException")) && !bytes.Contains(body, []byte("Faild"))
	return
}
