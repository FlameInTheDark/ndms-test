package api

import (
	"net/http"
	"strconv"
	"time"
)

func (api *API) register() {
	api.mux.HandleFunc("GET /{topic}", func(w http.ResponseWriter, r *http.Request) {
		topic := r.PathValue("topic")
		var timeout *time.Duration
		if t, ok := r.URL.Query()["timeout"]; ok {
			val, err := strconv.Atoi(t[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			td := time.Second * time.Duration(val)
			timeout = &td
		}
		data := api.topics.Consume(topic, timeout)
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(data)
	})

	api.mux.HandleFunc("POST /{topic}", func(w http.ResponseWriter, r *http.Request) {
		topic := r.PathValue("topic")

		v, ok := r.URL.Query()["v"]
		if !ok || len(v[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		api.topics.Produce(topic, []byte(v[0]))

		w.WriteHeader(http.StatusAccepted)
	})
}
