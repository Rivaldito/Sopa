package sopa

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html/charset"
)

func (sopa Sopa) Get(u string) {
	response, err := sopa.Client.Get(u)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	utf8Body, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(utf8Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
