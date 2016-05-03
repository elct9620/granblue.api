// Job Type
package granblue

type JobType int

const (
	UnknownJob = iota
	Attack
	Defense
	Heal
	Balance
	Special
)
