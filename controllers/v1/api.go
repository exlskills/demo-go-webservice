package v1

import (
	"github.com/exlinc/golang-utils/jsonhttp"
	"net/http"
)

func API(w http.ResponseWriter, r *http.Request) {
	// TODO check service health, send some metrics, etc.
	jsonhttp.JSONSuccess(w, nil, "Server healthy")
}
