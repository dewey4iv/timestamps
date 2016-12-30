package timestamps

import "time"

// Now retruns a pointer to the current time
func Now() *time.Time {
	t := time.Now()
	return &t
}
