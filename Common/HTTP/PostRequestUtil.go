package HTTP

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HTTP_Post_JSON_Payload(url string, headers map[string]string, body []byte) string {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	fmt.Println(req)

	handle_error(err)

	log.Println("Setting the request headers.")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	handle_error(err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		panic(resp.Status)
	}

	resp_body, err := io.ReadAll(resp.Body)
	handle_error(err)

	return string(resp_body)
}
