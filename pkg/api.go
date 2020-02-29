package api

import (
	"github.com/gorilla/mux"
)

func CreateApi() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getCCPosts", GetCCPosts)
	return r
}
