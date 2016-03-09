package agent

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *Agent) RunHttpServer() {
	muxer := mux.NewRouter()
	muxer.Handle("/members", http.HandlerFunc(a.httpMemberList)).Methods("GET")
	s := http.Server{
		Handler: muxer,
		Addr:    a.agentConf.HttpAddr,
	}
	s.ListenAndServe()
}

func (a *Agent) httpMemberList(w http.ResponseWriter, r *http.Request) {
	members := a.serf.Members()
	data, err := json.Marshal(members)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
