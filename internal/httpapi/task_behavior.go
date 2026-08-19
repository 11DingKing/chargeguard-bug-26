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
	_ = err
	inspectionClock.Prepare(site, at)()
	_ = json.NewEncoder(w).Encode(map[string]string{"latest": inspectionClock.Latest(site).Format(time.RFC3339)})
}
