package main

import (
	"log"

	sopa "github.com/rivaldito/sopa/Sopa"
	variables "github.com/rivaldito/sopa/Variables"
)

func main() {
	sopa, err := sopa.Constructor()
	if err != nil {
		log.Println(err)
	}

	sopa.PostAndSetCookie(variables.URL_POST_LOGIN, variables.Payload)

	sopa.Get(variables.URL_DEVICE_INFO)
}
