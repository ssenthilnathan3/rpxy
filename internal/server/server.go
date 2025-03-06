package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ssenthilnathan3/rpxy/internal/config"
)

func Run() error {
	config, err := config.NewConfiguration()
	if err != nil {
		return fmt.Errorf("error creating configuration %s", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)

	for _, resources := range config.Resources {
		url, _ := url.Parse(resources.Destination_url)
		proxy := NewProxy(*url)

		mux.HandleFunc(resources.Endpoint, ProxyRequestHandler(proxy, url, resources.Endpoint))
	}

	if err := http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("could not start the HTTP server %s", err)
	}

	return nil
}
