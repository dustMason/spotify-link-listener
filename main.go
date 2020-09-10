package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", Server)
	http.ListenAndServe(":1337", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p, err := ParseUri(r.Form.Get("link"))
	if err != nil {
		http.Error(w, "could not parse uri", 500)
		return
	}

	b, err := ioutil.ReadFile("/home/pi/RPi-Jukebox-RFID/settings/Latest_RFID")
	if err != nil {
		http.Error(w, "could not read latest rfid", 500)
		return
	}

	cardID := strings.TrimSuffix(string(b), "\n")

	form := url.Values{
		"cardID":         {cardID},
		"streamURL":      {p},
		"audiofolderNew": {strings.ReplaceAll(p, ":", "-")},
		"TriggerCommand": {"false"},
		"streamType":     {"spotify"},
		"submit":         {"submit"},
	}
	resp, err := http.PostForm("http://localhost:80/cardRegisterNew.php", form)
	if err != nil {
		http.Error(w, "failed to post", 500)
		return
	}
	if resp.StatusCode != 200 {
		http.Error(w, "got non-200", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	errMsg := GetError(resp.Body)
	if errMsg != "" {
		fmt.Println(errMsg)
		http.Error(w, errMsg, 500)
		return
	}

	fmt.Printf("associated %v with %v\n", cardID, p)
	fmt.Fprint(w, "")
}

func ParseUri(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	pathParts := strings.Split(u.Path[1:], "/")
	return "spotify:" + strings.Join(pathParts, ":"), nil
}

func GetError(r io.Reader) string {
	t := html.NewTokenizer(r)
	for {
		tt := t.Next()
		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.StartTagToken:
			to := t.Token()
			if to.Data == "div" {
				for _, a := range to.Attr {
					if a.Key == "class" && a.Val == "alert alert-danger" {
						var out string
						for {
							tn := t.Next()
							tnt := t.Token()
							if tn == html.TextToken {
								out = out + tnt.String()
							}
							if tn == html.EndTagToken && tnt.Data == "div" {
								return out
							}
						}
					}
				}
			}
		}
	}
}
