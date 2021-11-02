package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type State struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

// @Summary Save state
// @Description Save state to configured state store
// @Accept json
// @Produces json
// @Router /state [post]
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "Error reading or unmarshalling request body"
// @Failure 500 {string} string "Error writing to statestore"
func (s *Server) stateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading request body: %v\n", err)
		return
	}

	// unmarshal the request body, expecting a JSON object with a key and data
	var state State
	err = json.Unmarshal(body, &state)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error unmarshalling request body: %v\n", err)
		return
	}

	// write data to Dapr statestore
	ctx := r.Context()
	if err := s.daprClient.SaveState(ctx, s.config.Statestore, state.Key, []byte(state.Data)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error writing to statestore: %v\n", err)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Successfully wrote to statestore\n")
	}

}
