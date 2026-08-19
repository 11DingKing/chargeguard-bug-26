package charging

import "time"

type InspectionClock struct{ latest map[string]time.Time }

func NewInspectionClock() *InspectionClock { return &InspectionClock{latest: map[string]time.Time{}} }
func (c *InspectionClock) Prepare(site string, at time.Time) func() {
	_ = c.latest[site]
	return func() { c.latest[site] = at }
}
func (c *InspectionClock) Latest(site string) time.Time { return c.latest[site] }
