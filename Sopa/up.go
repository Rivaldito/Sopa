package sopa

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (sopa Sopa) Up(URL_DATA string) {

	var buf bytes.Buffer

	writer := multipart.NewWriter(&buf)

	file, err := os.Open("up3.img")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer.FormDataContentType()

	fw, err := writer.CreateFormFile("file", "up3.img")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(fw, file); err != nil {
		log.Fatal(err)
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, URL_DATA, &buf)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Type", "boundary=---------------------------87606390823769235343876761374")

	resp, err := sopa.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

func (sopa Sopa) Up2(URL_FW string, data map[string]string, fileData []byte) {

	boundary := "87606390823769235343876761374"

	var buffer bytes.Buffer
	// Add header for "onSubmit" field
	buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buffer.WriteString("Content-Disposition: form-data; name=\"onSubmit\"\r\n")
	buffer.WriteString("\r\n")
	buffer.WriteString(data["onSubmit"] + "\r\n")
	// Add header for "sessionkey" field
	buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buffer.WriteString("Content-Disposition: form-data; name=\"sessionkey\"\r\n")
	buffer.WriteString("\r\n")
	buffer.WriteString(data["sessionkey"] + "\r\n")
	// Add header for file upload
	buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buffer.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"FW_File\"; filename=\"%s\"\r\n", "up3.img"))
	buffer.WriteString("Content-Type: application/octet-stream\r\n")
	buffer.WriteString("\r\n")

	buffer.Write(fileData)
	buffer.WriteString("\r\n")

	// Add closing boundary
	buffer.WriteString(fmt.Sprintf("--%s--\r\n", boundary))

	// Create a POST request with the provided URL and boundary in the Content-Type header
	req, err := http.NewRequest("POST", URL_FW, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		fmt.Println("PRIMER ERROR")
		log.Println(err)
	}
	//req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	req.Header.Set("Content-Type", "multipart/form-data;boundary=---------------------------87606390823769235343876761374")
	resp, err := sopa.Client.Do(req)
	if err != nil {
		fmt.Println("Segundo ERROR")
		log.Println(err)
	}

	// Handle the response (optional)
	defer resp.Body.Close()
}
