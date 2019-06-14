package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getAntiForgeryField(doc *html.Node) string {
	var value string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "form" && n.Data == "input" {
			for _, attr := range n.Attr {
				if attr.Key == "value" {
					value = attr.Val
					return
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

func createSession() (cookie string, sessionID string, applicationCookie string) {
	res, _ := http.Get("http://dormitory.sutech.ac.ir/Login/Login")
	cookie = strings.Split(strings.Split(res.Header.Get("Set-Cookie"), ";")[0], "=")[1]
	doc, _ := html.Parse(res.Body)
	q := request{
		cookie:      cookie,
		antiForgery: getAntiForgeryField(doc),
		username:    "95113042",
		password:    "Hydro7790",
	}

	body := strings.NewReader(fmt.Sprintf("-----------------------------1341080666405880606225559037\r\nContent-Disposition: form-data; name=\"__RequestVerificationToken\"\r\n\r\n%s\r\n-----------------------------1341080666405880606225559037\r\nContent-Disposition: form-data; name=\"Entity.UserName\"\r\n\r\n%s\r\n-----------------------------1341080666405880606225559037\r\nContent-Disposition: form-data;  name=\"Entity.Password\"\r\n\r\n%s\r\n-----------------------------1341080666405880606225559037--", q.antiForgery, q.username, q.password))
	// fmt.Println(len("1341080666405880606225559037"))

	req, _ := http.NewRequest("POST", "http://dormitory.sutech.ac.ir/Login/Login", body)
	req.Header.Add("Host", "dormitory.sutech.ac.ir")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:66.0) Gecko/20100101 Firefox/66.0")
	// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	// req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Referer", "http://dormitory.sutech.ac.ir")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=---------------------------1341080666405880606225559037")
	// req.Header.Add("Content-Length", "554")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", fmt.Sprintf("__RequestVerificationToken=%s", q.cookie))
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("DNT", "1")

	res, _ = client.Do(req)

	// f, _ := os.Create("response.html")
	// res.Write(f)
	cookies := res.Cookies()
	sessionID = cookies[0].Value
	applicationCookie = cookies[1].Value
	registerSession(cookie, sessionID, applicationCookie)
	return
}
