package goravel

import (
	"encoding/json"
	"net/http"
)

func (grv *Goravel) WriteJSON(rw http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			rw.Header()[key] = val
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, err = rw.Write(out)
	if err != nil {
		return err
	}

	return nil
}
