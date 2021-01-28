package http

import (
	"encoding/json"
	"net/http"
)

func SendResponse(res interface{}, w http.ResponseWriter, err error) error {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	bRes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshaling response"))
		return err
	}

	_, err = w.Write(bRes)
	if err != nil {
		return err
	}
	return nil
}
