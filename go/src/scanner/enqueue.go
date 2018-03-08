package scanner

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	// "log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Enqueue(filename string) (apiResponse *EnqueueResponse, err error) {
	apiResponse = &EnqueueResponse{}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	postBody := &bytes.Buffer{}
	writer := multipart.NewWriter(postBody)
	part, err := writer.CreateFormFile("archive", filepath.Base(filename))
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	if err = writer.Close(); err != nil {
		return
	}

	request, err := http.NewRequest("POST", api_enqueue, postBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, apiResponse)
	return
}
