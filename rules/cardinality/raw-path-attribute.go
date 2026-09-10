package handler

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
)

func recordBad(r *http.Request) {
	// ruleid: tqa-card-raw-path-attribute-go
	attribute.String("http.target", r.URL.Path)

	// ruleid: tqa-card-raw-path-attribute-go
	attribute.String("http.uri", r.RequestURI)
}

func recordGood(routePattern string) {
	// ok: tqa-card-raw-path-attribute-go
	attribute.String("http.route", routePattern)
}