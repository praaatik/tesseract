package worker

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
	"log/slog"
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

	// Logger is used to log the relevant information/error/warning messages.
	Logger *slog.Logger
}

// Log logs a message with the specified level and arguments.
func (w *Worker) Log(prefix string, level slog.Level, message string, args ...any) {
	switch level {
	case slog.LevelInfo:
		w.Logger.Info(fmt.Sprintf("%s running", prefix), "worker_name", w.Name, "task_count", w.TaskCount, "task_queue", w.TaskQueue, "task_db", w.TaskDb, "additional_info", args)
	default:
		w.Logger.Error(fmt.Sprintf("%s running", prefix), "worker_name", w.Name, "task_count", w.TaskCount, "task_queue", w.TaskQueue, "task_db", w.TaskDb, "additional_info", args)
	}
}

func (w *Worker) StartTask() {
	w.Log("StartTask", slog.LevelInfo, "Starting a task", "reason", "user-requested")
}

func (w *Worker) StopTask() {
	w.Log("StopTask", slog.LevelInfo, "Stopping a task", "reason", "user-requested")
}

func (w *Worker) RunTask() {
	w.Log("RunTask", slog.LevelInfo, "Running a task", "reason", "user-requested")
}

func (w *Worker) CollectStatistics() {
	w.Log("CollectStatistics", slog.LevelInfo, "Collecting stats", "reason", "automated")
}
