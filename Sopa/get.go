package sopa

import (
	"io"
	"log"

	"golang.org/x/net/html/charset"
)

func (sopa *Sopa) Get(u string) error {
	response, err := sopa.Client.Get(u)
	if err != nil {
		log.Panic(err)
		return err
	}
	defer response.Body.Close()

	utf8Body, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		panic(err)
		//return nil, err
	}
	bytes, err := io.ReadAll(utf8Body)
	if err != nil {
		panic(err)
		//return nil, err
	}
	html := string(bytes)

	sopa.HTMLResponse = &html

	return nil
}
