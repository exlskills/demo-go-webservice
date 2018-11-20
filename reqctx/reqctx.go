package reqctx

import (
	"context"
	"errors"
	"github.com/exlinc/golang-utils/htmlhttp"
	"github.com/exlinc/golang-utils/jsonhttp"
	"github.com/exlskills/demo-go-webservice/models"
	"net/http"
)

const currentUserContextKey = "req_usr"

var unauthorizedError = jsonhttp.APIResponse{Message: "Unauthorized", Success: false, Debug: "Token/account error"}

func AddCurrentUserToContext(r *http.Request, user models.User) context.Context {
	return context.WithValue(r.Context(), currentUserContextKey, user)
}

func GetCurrentUser(r *http.Request) (models.User, error) {
	user := r.Context().Value(currentUserContextKey)
	switch user.(type) {
	case nil:
		return models.User{}, errors.New("current user not found")
	default:
		return user.(models.User), nil
	}
}

func GetCurrentUserAndCatchForAPI(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	user, err := GetCurrentUser(r)
	if err != nil || user.ID == 0 {
		jsonhttp.JSONWriter(w, unauthorizedError, http.StatusUnauthorized)
		return user, err
	}
	return user, nil
}

func GetCurrentUserAndCatchForView(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	user, err := GetCurrentUser(r)
	if err != nil || user.ID == 0 {
		htmlhttp.UnauthorizedErrorView(w, r)
		return user, err
	}
	return user, nil
}
