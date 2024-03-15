package sopa

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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

}

func (sopa Sopa) Post(u string, payload url.Values) {

	response, err := sopa.Client.PostForm(u, payload)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

}

func (sopa Sopa) PostUploadBinaryFile(u string, payload url.Values, upload []byte) {

	//reader := bytes.NewReader(upload)

	request, err := http.NewRequest("POST", variables.URL_FW, strings.NewReader(payload.Encode()))
	if err != nil {
		log.Panic(err)
	}
	//request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Content-Type", "application/octet-stream")

	rsp, err := sopa.Client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}

	sopa.Client.Do(request)

}
