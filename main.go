package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

var file *os.File
var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// if req.Method == http.MethodGet && req.URL.Path == "/" && via[0].Method != http.MethodPost {
		// 	return nil
		// }
		return errors.New("We Hate redirects! :)")
	},
}

func init() {
	file, _ = os.Create("response.html")
}

type reserveRequest struct {
	RoomID             string `json:"roomId"`
	ReserveDescription string `json:"reserveDescription"`
	TermID             string `json:"termId"`
}

var defaultResReq = &reserveRequest{
	RoomID:             "5d29855d-3924-e811-80c2-005056aa18ec",
	ReserveDescription: "خوابگاه مدرس، بلوک 1، طبقه 1، اتاق شماره 116",
}

func (r *reserveRequest) fillTermID(cookie, sessionID, applicationCookie string) {
	req, _ := http.NewRequest("GET", "http://dormitory.sutech.ac.ir/Dorm/RoomReserve", nil)
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
	res, _ := client.Do(req)
	// res.Write(file)
	doc, _ := html.Parse(res.Body)
	r.TermID = getTermID(doc)
}

type request struct {
	cookie      string
	antiForgery string
	username    string
	password    string
}

func main() {
	var (
		count1, count2                       int
		cookie, sessionID, applicationCookie string
	)
	cookie, sessionID, applicationCookie = createSession()
	for defaultResReq.TermID == "" {
		if count1 > 30 {
			cookie, sessionID, applicationCookie = createSession()
			count1 = 0
		}
		defaultResReq.fillTermID(cookie, sessionID, applicationCookie)
		fmt.Println("termId:", defaultResReq.TermID)
		count1++
		time.Sleep(time.Second * 2)
	}

	for count2 < 10 {
		reserve(cookie, sessionID, applicationCookie)
		count2++
		time.Sleep(time.Second * 1)
	}
}
