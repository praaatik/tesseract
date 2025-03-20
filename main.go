package main

import (
	"github.com/golang-collections/collections/queue"
	"github.com/praaatik/tesseract/logger"
	"github.com/praaatik/tesseract/worker"
)

func main() {
	w1 := worker.Worker{
		Name:      "w1",
		TaskQueue: queue.Queue{},
		TaskDb:    nil,
		TaskCount: 0,
		Logger:    logger.Logger.With("worker", "w1"),
	}
	w1.StartTask()

	w2 := worker.Worker{
		Name:      "w2",
		TaskQueue: queue.Queue{},
		TaskDb:    nil,
		TaskCount: 0,
		Logger:    logger.Logger.With("worker", "w2"),
	}
	w2.StartTask()
	w2.CollectStatistics()
}
