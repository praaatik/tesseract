package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/praaatik/tesseract/node"
	"github.com/praaatik/tesseract/task"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	e := task.Event{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}
	fmt.Println(e)

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}
	fmt.Printf("node: %v\n", n)
}
