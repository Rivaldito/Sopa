package sopa

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Sopa struct {
	Client       *http.Client
	HTMLResponse *string
}

func Constructor() (*Sopa, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{
		Timeout: time.Duration(160) * time.Second,
		Jar:     jar,
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	sopa := Sopa{
		Client: client,
	}

	return &sopa, nil
}

func (sopa Sopa) Test() {
	fmt.Println(sopa.Client)
}
