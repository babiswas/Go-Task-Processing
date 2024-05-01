package HTTP

import (
	"io"
	"log"
	"net/http"
)

func HTTP_Get(url string, headers map[string]string) string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	handle_error(err)

	log.Println("Setting the request headers.")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	handle_error(err)

	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	handle_error(err)
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	return string(resp_body)

}
