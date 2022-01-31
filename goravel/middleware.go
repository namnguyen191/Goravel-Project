package goravel

import "net/http"

func (grv *Goravel) SessionLoad(next http.Handler) http.Handler {
	return grv.Session.LoadAndSave(next)
}
