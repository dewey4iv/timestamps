package timestamps

import "time"

// Option applies changes to the timestamp
type Option interface {
	Apply(*Timestamps)
}

// Updated marks a timestamp set updated
var Updated *update

type update struct{}

func (opt *update) Apply(t *Timestamps) {
	t.UpdatedAt = Now()
}

// Created marks a timestamp as created
var Created *created

type created struct{}

func (opt *created) Apply(t *Timestamps) {
	now := Now()
	t.UpdatedAt = now
	t.CreatedAt = now
}

// UpdatedWithTime sets the timestamps as updated with the provided time
func UpdatedWithTime(updated time.Time) Option {
	return &updatedWithTime{updated}
}

type updatedWithTime struct {
	updated time.Time
}

func (opt *updatedWithTime) Apply(t *Timestamps) {
	t.Updated = &opt.updated
}

// CreatedWithTime sets the timestamps as created with the provided time
func CreatedWithTime(created time.Time) Option {
	return &createdWithTime{created}
}

type createdWithTime struct {
	created time.Time
}

func (opt *createdWithTime) Apply(t *Timestamps) {
	t.Created = &opt.created
	t.Updated = &opt.created
}
