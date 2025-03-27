package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/logger"
	"github.com/praaatik/tesseract/task"
	"github.com/praaatik/tesseract/worker"
)

func main() {

	logLevelStr := flag.String("loglevel", "INFO", "Set logging level (DEBUG, INFO, WARN, ERROR)")
	flag.Parse()

	logLevel, err := logger.ParseLevel(*logLevelStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid log level %q, defaulting to INFO: %v\n", *logLevelStr, err)
		logLevel = logger.INFO
	}

	db := make(map[uuid.UUID]*task.Task)
	logger := logger.NewLogger("main: ", logLevel)

	w := worker.Worker{
		TaskQueue: queue.New(),
		TaskDb:    db,
		Logger:    logger,
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
