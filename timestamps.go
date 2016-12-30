package timestamps

import "time"

// Timestamps holds the created and updated fields of a timestamp
type Timestamps struct {
	Updated *time.Time `json:"updated" gorethink:"updated,omitempty"`
	Created *time.Time `json:"created" gorethink:"created,omitempty"`
}

// Mark takes a variadic list of Options and applies them to the timestamp
func (ts *Timestamps) Mark(opts ...Option) {
	if len(opts) == 0 {
		t := Now()

		ts.Created = t
		ts.Updated = t

		return
	}

	for _, opt := range opts {
		opt.Apply(ts)
	}
}
