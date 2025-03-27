// Worker responsibilities:
// 1. Run tasks as Docker containers
// 2. Accept tasks to run from a manager
// 3. Provide relevant statistics to the manager for the purpose of scheduling tasks
// 4. Keep track of its tasks and their state

package worker

import (
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/logger"
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

	Logger *logger.Logger
}

func (w *Worker) StartTask(t task.Task) task.DockerResult {
	t.StartTime = time.Now().UTC()
	w.Logger.Info("Starting task: %v", t.ID)

	config := task.NewConfig(&t)
	d := task.NewDocker(config)

	result := d.Run()
	if result.Error != nil {
		w.Logger.Error("Error running task %v: %v", t.ID, result.Error)
		t.State = task.Failed
		w.TaskDb[t.ID] = &t

		return result
	}

	t.ContainerID = result.ContainerId
	t.State = task.Running
	w.TaskDb[t.ID] = &t

	return result
}

func (w *Worker) AddTask(t task.Task) {
	w.TaskQueue.Enqueue(t)
	w.Logger.Debug("Task %v added to the queue TaskQueue", t.ID)
}

func (w *Worker) StopTask(t task.Task) task.DockerResult {
	w.Logger.Info("Stopping task %v with container %v", t.ID, t.ContainerID)

	config := task.NewConfig(&t)
	d := task.NewDocker(config)

	result := d.Stop(t.ContainerID)

	if result.Error != nil {
		w.Logger.Error("Error stopping container %v: %v", t.ContainerID, result.Error)
	}

	t.FinishTime = time.Now().UTC()
	t.State = task.Completed
	w.TaskDb[t.ID] = &t
	// log.Printf("Stopped and removed container %v for task %v\n",
	// 	t.ContainerID, t.ID)

	w.Logger.Info("Stopped and removed container %v for task %v", t.ContainerID, t.ID)

	return result
}

func (w *Worker) RunTask() task.DockerResult {
	t := w.TaskQueue.Dequeue()
	if t == nil {
		w.Logger.Warn("No tasks in the queue")
		return task.DockerResult{Error: nil}
	}
	taskQueued := t.(task.Task)
	taskPersisted := w.TaskDb[taskQueued.ID]
	if taskPersisted == nil {
		taskPersisted = &taskQueued
		w.TaskDb[taskQueued.ID] = taskPersisted
	}

	var result task.DockerResult
	if task.ValidStateTransition(taskPersisted.State, taskQueued.State) {
		switch taskQueued.State {
		case task.Scheduled:
			result = w.StartTask(taskQueued)
		case task.Completed:
			result = w.StopTask(taskQueued)
		default:
			w.Logger.Error("Invalid state encountered: %v", result.Error)

		}
	} else {
		err := fmt.Errorf("Invalid transition from %v to %v", taskPersisted.State, taskQueued.State)
		w.Logger.Warn("Invalid state transition: %v", err)

		result.Error = err
		return result
	}

	return result
}

func (w *Worker) CollectStatistics() {
	// TODO: implementation
	w.Logger.Debug("Collecting statistics (not implemented yet).")
}
