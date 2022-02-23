package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")
	if len(domainParts) > 1 {
		subdomain := domainParts[0]
		if handler, ok := subdomains[subdomain]; ok {
			handler.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func loadHandlers(config *SimpleProxy) Subdomains {
	subdomains := make(Subdomains)
	for _, route := range config.Routes {
		if route.Subdomain != nil {
			subdomains[*route.Subdomain] = handleRoute(route)
			log.Printf("Route %s added to subdomain %s\n", route.To, *route.Subdomain)
		} else {
			log.Printf(`could not find attribute "subdomain" for route with name: "%s"`, route.Name)
		}
	}
	return subdomains
}

func handleRoute(route Route) http.Handler {
	remote, err := url.Parse(route.To)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	return proxy
}
