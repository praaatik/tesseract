package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
	"github.com/praaatik/tesseract/worker"
)

func main() {
	db := make(map[uuid.UUID]*task.Task)
	w := worker.Worker{
		TaskQueue: queue.New(),
		TaskDb:    db,
	}

	t := task.Task{
		ID:    uuid.New(),
		Name:  fmt.Sprintf("test-container-1-%d", rand.IntN(1000)),
		State: task.Scheduled,
		Image: "strm/helloworld-http",
	}

	// first time the worker will see the task
	w.AddTask(t)
	result := w.RunTask()
	if result.Error != nil {
		panic(result.Error)
	}

	t.ContainerID = result.ContainerId

	time.Sleep(time.Second * 1)

	t.State = task.Completed
	w.AddTask(t)
	result = w.RunTask()
	if result.Error != nil {
		panic(result.Error)
	}
}
