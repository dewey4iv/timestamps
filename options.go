package timestamps

// Option applies changes to the timestamp
type Option interface {
	Apply(*Timestamp)
}

// Updated marks a timestamp set updated
var Updated = func() Option {
	return &update{}
}

type update struct{}

func (opt *update) Apply(t *Timestamp) {
	t.Updated = Now()
}

// Created marks a timestamp as created
var Created = func() Option {
	return &created{}
}

type created struct{}

func (opt *created) Apply(t *Timestamp) {
	now := Now()
	t.Updated = now
	t.Created = now
}
