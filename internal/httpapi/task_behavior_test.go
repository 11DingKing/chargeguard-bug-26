package httpapi

import (
	"chargeguard/internal/charging"
	"testing"
	"time"
)

func TestTaskBehavior(t *testing.T) {
	clock := charging.NewInspectionClock()
	old, _ := time.Parse(time.RFC3339, "2026-03-01T08:00:00Z")
	newer, _ := time.Parse(time.RFC3339, "2026-03-01T09:00:00Z")
	commitOld := clock.Prepare("station-1", old)
	commitNew := clock.Prepare("station-1", newer)
	done := make(chan struct{})
	go func() { commitNew(); close(done) }()
	<-done
	goDone := make(chan struct{})
	go func() { commitOld(); close(goDone) }()
	<-goDone
	if got := clock.Latest("station-1"); !got.Equal(newer) {
		t.Fatalf("latest=%s", got)
	}
}
