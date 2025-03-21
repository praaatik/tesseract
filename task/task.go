package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"time"
)

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

// Task is the smallest unit of work to be performed.
// This struct might be extended with more fields.
type Task struct {
	// ID is a unique identifier field for the Task.
	ID uuid.UUID

	// Human-readable name format of the Task.
	Name string

	// State is the current lifecycle state of the Task
	State State

	// Image indicates the Docker image the Task is running.
	Image string

	// Memory is useful to identify the memory the Task would require.
	Memory int

	// Disk is useful to identify the disk the Task would require.
	Disk int

	// ExposedPorts defines the ports that the Task container will expose.
	ExposedPorts nat.PortSet

	// PortBindings maps container ports to host port.
	PortBindings map[string]string

	// RestartPolicy defines the policy which tells the system what to do when a Task fails.
	RestartPolicy string

	// StartTime is the time when the Task was started.
	StartTime time.Time

	// FinishTime is the time when the Task was completed.
	FinishTime time.Time
}

// Event represents a change in the Task state.
// Users don't interact with this, it is triggered whenever there is a change in the Task state.
// Renamed from TaskEvent to Event because of - https://go.dev/blog/package-names
type Event struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
