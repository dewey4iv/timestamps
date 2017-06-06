package timestamps

import "time"

// Timestamps holds the created and updated fields of a timestamp
type Timestamps struct {
	UpdatedAt *time.Time `json:"updated_at" gorethink:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at" gorethink:"created_at,omitempty"`
}

// Mark takes a variadic list of Options and applies them to the timestamp
func (ts *Timestamps) Mark(opts ...Option) {
	if len(opts) == 0 {
		t := Now()

		ts.CreatedAt = t
		ts.UpdatedAt = t

		return
	}

	for _, opt := range opts {
		opt.Apply(ts)
	}
}
