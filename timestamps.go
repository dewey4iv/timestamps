package timestamps

import (
	"time"

	"github.com/the-control-group/watchdog-api/util"
)

type Timestamps struct {
	Updated *time.Time `json:"updated" gorethink:"updated,omitempty"`
	Created *time.Time `json:"created" gorethink:"created,omitempty"`
}

func (ts *Timestamps) MarkCreated(t *time.Time) {
	if t == nil {
		t = util.Now()
	}

	ts.Created = t
	ts.Updated = t
}

/*
	Passing a func() bool allows for more complicated logic
	AND maintains reability.
	Usage:
	```
		ts := new(Timestamps)
		ts.MarkUpdated(nil, AndNilCreated)
	```
*/

func (ts *Timestamps) MarkUpdated(t *time.Time, resetCreated func() bool) {
	if t == nil {
		t = util.Now()
	}

	if resetCreated == nil || resetCreated() {
		ts.Created = nil
	}

	ts.Updated = t
}

func AndNilCreated() bool {
	return true
}
