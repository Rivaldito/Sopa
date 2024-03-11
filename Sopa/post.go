package sopa

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	variables "github.com/rivaldito/sopa/Variables"
)

func (sopa Sopa) PostAndSetCookie(u string, payload url.Values) {

	response, err := http.PostForm(u, payload)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	uParse, err := url.Parse(variables.URL)
	if err != nil {
		log.Panic(err)
	}

	sopa.Client.Jar.SetCookies(uParse, response.Cookies())

	fmt.Println(sopa.Client.Jar)

}

func (sopa Sopa) Post(u string, payload url.Values) {

	response, err := http.PostForm(u, payload)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

}
