// Worker responsibilities:
// 1. Run tasks as Docker containers
// 2. Accept tasks to run from a manager
// 3. Provide relevant statistics to the manager for the purpose of scheduling tasks
// 4. Keep track of its tasks and their state

package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
)

type Worker struct {
	// Name is the human-readable name of the Worker
	Name string

	// TaskQueue is queue to accept Task from the manager and execute in a FIFO manner.
	TaskQueue *queue.Queue

	// TaskDb keeps a track of the Task and it's state.
	TaskDb map[uuid.UUID]*task.Task

	// TaskCount keeps a track of the number of Tasks at any given time.
	TaskCount int
}

func (w *Worker) StartTask() {
	//TODO: implementation
}

func (w *Worker) StopTask() {
	//TODO: implementation
}

func (w *Worker) RunTask() {
	// TODO: implementation
}

func (w *Worker) CollectStatistics() {
	// TODO: implementation
}
