package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ProxyHandler(location string, proxyUrl string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiUrl := proxyUrl + strings.TrimPrefix(r.URL.Path, location)

		request, error := http.NewRequest(r.Method, apiUrl, r.Body)

		if error != nil {
			fmt.Println(error)
		}

		client := &http.Client{}
		response, error := client.Do(request)

		if error != nil {
			fmt.Println(error)
		}

		responseBody, error := io.ReadAll(response.Body)

		if error != nil {
			fmt.Println(error)
		}

		w.Write(responseBody)

		defer response.Body.Close()
	}
}
