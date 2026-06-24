package api

import (
	"net/http"

	"github.com/FlameInTheDark/ndms-test/internal/topics"
)

type API struct {
	topics *topics.Topics
	mux    *http.ServeMux
}

func NewAPI() *API {
	return &API{
		topics: topics.NewTopics(),
		mux:    http.NewServeMux(),
	}
}

func (api *API) Serve(baseUrl string) error {
	api.register()
	return http.ListenAndServe(baseUrl, api.mux)
}
