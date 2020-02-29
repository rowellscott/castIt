package api

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCCPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.facebook.com/centralcastinglosangeles/")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request to Central Casting Los Angeles Failed"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading response body from Central Casting Los Angeles"))
	}
	log.Info(body)
}
