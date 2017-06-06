package timestamps

import "time"

func init() {
	time.Local = time.UTC
}

// Now retruns a pointer to the current time
func Now() *time.Time {
	t := time.Now().UTC()
	return &t
}
