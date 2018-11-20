package v1

import (
	"github.com/exlinc/golang-utils/jsonhttp"
	"github.com/exlinc/golang-utils/queryparams"
	"github.com/exlskills/demo-go-webservice/models"
	"net/http"
)

func GetGophers(w http.ResponseWriter, r *http.Request) {
	limit, offset := queryparams.GetLimitOffsetQueryParametersDefaults(r)

	gophers, err := models.GetGophers(limit, offset)
	if err != nil {
		jsonhttp.JSONNotFoundError(w, "Error fetching gophers", "")
		return
	}

	jsonhttp.JSONSuccess(w, map[string]interface{}{"gophers": gophers}, "Successfully queried gophers")
}
