// Manager responsibilities:
// 1. Accept requests from users to start and stop tasks
// 2. Schedule tasks onto worker machines
// 3. Keep track of tasks, their states, and the machine on which they run

package manager

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
)

type Manager struct {
	// Pending is a queue having the Task which are in the pending state of their lifecycle.
	Pending queue.Queue

	// TaskDb stores the tasks
	TaskDb map[string][]*task.Task

	// EventDb stores the events
	EventDb map[string][]*task.Event

	// Workers will keep a track of all the workers which are currently running Tasks.
	Workers []string

	// TODO: check the use of WorkerTaskMap, not entirely sure yet
	WorkerTaskMap map[string][]uuid.UUID

	// TODO: check the use of TaskWorkerMap , not entirely sure yet
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
}

func (m *Manager) UpdateTasks() {
}

func (m *Manager) SendWork() {
}
