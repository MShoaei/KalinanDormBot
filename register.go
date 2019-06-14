package main

import (
	"fmt"
	"net/http"
)

func registerSession(cookie, sessionID, applicationCookie string) {
	req, _ := http.NewRequest("GET", "http://dormitory.sutech.ac.ir/", nil)
	req.Header.Add("Host", "dormitory.sutech.ac.ir")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Connection", "close")
	req.Header.Add("Cookie", fmt.Sprintf(`__RequestVerificationToken=%s; ASP.NET_SessionId=%s; .AspNet.ApplicationCookie=%s`, cookie, sessionID, applicationCookie))
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("DNT", "1")
	// req.Header.Add("Cache-Control", "max-age=0")
	_, _ = client.Do(req)
}
