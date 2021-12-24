package goravel

import "net/http"

func (grv *Goravel) SessionLoad(next http.Handler) http.Handler {
	grv.InfoLog.Println("SessionLoad called")
	return grv.Session.LoadAndSave(next)
}
