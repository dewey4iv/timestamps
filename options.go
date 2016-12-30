package timestamps

// Option applies changes to the timestamp
type Option interface {
	Apply(*Timestamps)
}

// Updated marks a timestamp set updated
var Updated *update

type update struct{}

func (opt *update) Apply(t *Timestamps) {
	t.Updated = Now()
}

// Created marks a timestamp as created
var Created *created

type created struct{}

func (opt *created) Apply(t *Timestamps) {
	now := Now()
	t.Updated = now
	t.Created = now
}
