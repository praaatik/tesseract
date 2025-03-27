package node

import "github.com/praaatik/tesseract/logger"

// Node represents the machine on which the Task is running.
type Node struct {
	// Name of the node
	Name string

	// Ip is the IP address which manager requires to send tasks to Nodes.
	Ip string

	Cores int

	// Memory is the maximum amount of Memory a Task can use.
	Memory int

	// MemoryAllocated is the amount of Memory which is currently being used by the Node executing a Task.
	MemoryAllocated int

	// Disk is the maximum amount of Disk a Task can use.
	Disk int

	// DiskAllocated is the amount of Disk which is currently being used by the Node executing a Task.
	DiskAllocated int

	// TODO: Check what does this do?
	Role string

	// TaskCount is the number of Tasks the Node uses to keep track.
	TaskCount int
	Logger    *logger.Logger
}
