package middleware

import (
	"github.com/exlinc/golang-utils/htmlhttp"
	"github.com/exlinc/golang-utils/jsonhttp"
	"github.com/exlskills/demo-go-webservice/config"
	"github.com/exlskills/demo-go-webservice/models"
	"github.com/exlskills/demo-go-webservice/reqctx"
	"net/http"
	"strings"
)

var Log = config.Cfg().GetLogger()

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}

func RecoverAndLog(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				jsonhttp.JSONInternalError(w, "An internal server error occurred", "Please try again in a few seconds")
				Log.Error("Panic occurred in HTTP handler:", r)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}

func GetContext(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var passItOn = func() {
			handler.ServeHTTP(w, r)
		}
		var handleNonEmptyToken = func(value string) {
			user := models.User{}
			// TODO Here we would actually use the token value, i.e., a JWT, to track down and verify our user!
			//      Error ignored since we're in TODO mode here. IRL you should check that!
			user.FindByID()
			// Update our request
			r = r.WithContext(reqctx.AddCurrentUserToContext(r, user))
			// Move on to the next middleware
			passItOn()
		}

		if c, err := r.Cookie(config.Cfg().TokenCookieName); err == nil {
			if c.Value != "" {
				handleNonEmptyToken(c.Value)
			} else {
				passItOn()
			}
		} else if ah := r.Header.Get("Authorization"); ah != "" {
			if len(ah) > 6 && strings.ToUpper(ah[0:7]) == "BEARER " {
				val := ah[7:]
				if val != "" {
					handleNonEmptyToken(val)
				} else {
					passItOn()
				}
			} else {
				passItOn()
			}
		} else {
			passItOn()
		}
	}
}

func RequireAPIKey(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			jsonhttp.JSONDetailed(w, jsonhttp.APIResponse{Message: "Unauthorized", Debug: "Invalid or missing API Key header/query parameter"}, http.StatusUnauthorized)
		}

		if r.URL.Query().Get("apiKey") != config.Cfg().ServiceAPIKey && r.Header.Get("x-api-key") != config.Cfg().ServiceAPIKey {
			respondUnauthorized()
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func RequireUserForAPI(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			jsonhttp.JSONDetailed(w, jsonhttp.APIResponse{Message: "Unauthorized", Debug: "Invalid or missing access token header/cookie"}, http.StatusUnauthorized)
		}
		user, err := reqctx.GetCurrentUser(r)
		if err != nil || user.ID == 0 {
			respondUnauthorized()
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}

func RequireUserForView(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			htmlhttp.UnauthorizedErrorView(w, r)
		}
		user, err := reqctx.GetCurrentUser(r)
		if err != nil || user.ID == 0 {
			respondUnauthorized()
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}
