package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
	"log"
)

type Worker struct {
	// Name is the human-readable name of the Worker
	Name string

	// TaskQueue is queue to accept Task from the manager and execute in a FIFO manner.
	TaskQueue queue.Queue

	// TaskDb keeps a track of the Task and it's state.
	TaskDb map[uuid.UUID]*task.Task

	// TaskCount keeps a track of the number of Tasks at any given time.
	TaskCount int

	Logger *log.Logger
}

func (w *Worker) StartTask() {
}

func (w *Worker) StopTask() {
}

func (w *Worker) RunTask() {
}

func (w *Worker) CollectStatistics() {
}
