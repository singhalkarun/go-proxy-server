package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall"
)

func ProxyHandler(location string, proxyUrl string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiUrl := proxyUrl + strings.TrimPrefix(r.URL.Path, location)

		request, error := http.NewRequest(r.Method, apiUrl, r.Body)

		if error != nil {
			fmt.Println(error)
			w.Write([]byte(error.Error()))
			return
		}

		client := &http.Client{}
		response, error := client.Do(request)

		if error != nil {
			fmt.Println(error)
			if errors.Is(error, syscall.ECONNREFUSED) {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			} else {
				w.Write([]byte(error.Error()))
			}
			return
		}

		responseBody, error := io.ReadAll(response.Body)

		if error != nil {
			fmt.Println(error)
			w.Write([]byte(error.Error()))
		} else {
			w.Write(responseBody)
		}

		defer response.Body.Close()
	}
}
