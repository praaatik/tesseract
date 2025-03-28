package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/logger"
	"github.com/praaatik/tesseract/task"
	"github.com/praaatik/tesseract/worker"
)

func main() {
	host := os.Getenv("CUBE_HOST")
	port, _ := strconv.Atoi(os.Getenv("CUBE_PORT"))

	logLevelStr := flag.String("loglevel", "INFO", "Set logging level (DEBUG, INFO, WARN, ERROR)")
	flag.Parse()

	logLevel, err := logger.ParseLevel(*logLevelStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid log level %q, defaulting to INFO: %v\n", *logLevelStr, err)
		logLevel = logger.INFO
	}

	logger := logger.NewLogger("main: ", logLevel)

	w := worker.Worker{
		TaskQueue: queue.New(),
		TaskDb:    make(map[uuid.UUID]*task.Task),
		Logger:    logger,
	}

	api := worker.Api{
		Address: host,
		Port:    port,
		Worker:  &w,
	}

	go runTasks(&w)
	api.Start()
}

func runTasks(w *worker.Worker) {
	for {
		if w.TaskQueue.Len() != 0 {
			result := w.RunTask()
			if result.Error != nil {
				w.Logger.Error("Error running task: %v\n", result.Error)
			}
		} else {
			w.Logger.Info("No tasks to process currently.\n")
		}
		w.Logger.Info("Sleeping for 10 seconds.")
		time.Sleep(10 * time.Second)
	}
}
