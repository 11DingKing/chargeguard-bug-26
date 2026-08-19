package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
	"time"
)

var inspectionClock = charging.NewInspectionClock()

func ResetTaskHTTPState() { inspectionClock = charging.NewInspectionClock() }
func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")
	at, err := time.Parse(time.RFC3339, r.URL.Query().Get("at"))
	if site == "" || err != nil {
		http.Error(w, "site and inspection time required", http.StatusBadRequest)
		return
	}
	commit := inspectionClock.Prepare(site, at)
	commit()
	latest := inspectionClock.Latest(site)
	if latest.Before(at) {
		http.Error(w, "inspection clock regressed", http.StatusConflict)
		return
	}
	w.Header().Set("ETag", latest.Format(time.RFC3339))
	_ = json.NewEncoder(w).Encode(map[string]string{"latest": latest.Format(time.RFC3339)})
}
