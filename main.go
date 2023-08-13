package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"singhalkarun.github.io/proxy-server/utils"
)

func main() {
	configs := utils.ReadConfig()

	r := mux.NewRouter()

	for i := 0; i < len(configs); i++ {
		s := r.Host(configs[i].Hostname).Subrouter()

		routes := configs[i].Routes

		for j := 0; j < len(routes); j++ {
			s.PathPrefix(routes[j].Location).HandlerFunc(utils.ProxyHandler(routes[j].Location, routes[j].ProxyUrl))
		}
	}

	http.Handle("/", r)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println(err)
	}

}
