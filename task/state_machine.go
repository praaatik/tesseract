package task

import "slices"

// State represents the current lifecycle state of a Task.
type State int

const (
	// Pending is the initial state when the Task is enqueued.
	Pending State = iota

	// Scheduled state indicates that the system has determined a machine can run a task, but in process to send it.
	Scheduled

	// Running state indicates that task has been moved to the machine and is executing.
	Running

	// Completed state indicates when Task has completed successfully.
	Completed

	// Failed state indicates the Task has stopped working as expected or crashed.
	Failed
)

var stateTransitionMap = map[State][]State{
	Pending:   {Scheduled},
	Scheduled: {Scheduled, Running, Failed},
	Running:   {Running, Completed, Failed},
	Completed: {},
	Failed:    {},
}

func Contains(states []State, state State) bool {
	return slices.Contains(states, state)
}

func ValidStateTransition(src State, dst State) bool {
	return Contains(stateTransitionMap[src], dst)
}
