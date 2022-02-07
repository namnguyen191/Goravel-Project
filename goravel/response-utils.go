package goravel

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
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

func (grv *Goravel) WriteXML(rw http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			rw.Header()[key] = val
		}
	}

	rw.Header().Set("Content-Type", "application/xml")
	rw.WriteHeader(status)
	_, err = rw.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (grv *Goravel) DownloadFile(rw http.ResponseWriter, r *http.Request, pathToFile, fileName string) error {
	fp := path.Join(pathToFile, fileName)
	fileToServe := filepath.Clean(fp)
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	http.ServeFile(rw, r, fileToServe)

	return nil
}

func (grv *Goravel) Error404(rw http.ResponseWriter, r *http.Request) {
	grv.ErrorStatus(rw, http.StatusNotFound)
}

func (grv *Goravel) Error500(rw http.ResponseWriter, r *http.Request) {
	grv.ErrorStatus(rw, http.StatusInternalServerError)
}

func (grv *Goravel) ErrorUnauthorized(rw http.ResponseWriter, r *http.Request) {
	grv.ErrorStatus(rw, http.StatusUnauthorized)
}

func (grv *Goravel) ErrorForbidden(rw http.ResponseWriter, r *http.Request) {
	grv.ErrorStatus(rw, http.StatusForbidden)
}

func (grv *Goravel) ErrorStatus(rw http.ResponseWriter, status int) {
	http.Error(rw, http.StatusText(status), status)
}
