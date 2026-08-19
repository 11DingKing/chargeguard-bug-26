package charging

import (
	"sync"
	"time"
)

type InspectionClock struct {
	latest map[string]time.Time
	mu     sync.Mutex
}

func NewInspectionClock() *InspectionClock { return &InspectionClock{latest: map[string]time.Time{}} }
func (c *InspectionClock) Prepare(site string, at time.Time) func() {
	return func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if current := c.latest[site]; current.IsZero() || at.After(current) {
			c.latest[site] = at
		}
	}
}
func (c *InspectionClock) Latest(site string) time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.latest[site]
}
