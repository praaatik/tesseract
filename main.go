package main

import (
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
)

func createContainer() (*task.Docker, *task.DockerResult) {
	c := task.Config{
		Name:  "hello-world-container-2",
		Image: "hello-world",
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := task.Docker{
		Client: dc,
		Config: c,
	}

	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil, nil
	}

	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)
	return &d, &result
}

func stopContainer(d *task.Docker, containerId string) *task.DockerResult {
	result := d.Stop(containerId)
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil
	}

	fmt.Printf("Container %s has been stopped and removed\n", result.ContainerId)
	return &result
}

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	te := task.Event{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	dockerTask, createResult := createContainer()
	if createResult.Error != nil {
		fmt.Printf("%v\n", createResult.Error)
		os.Exit(1)
	}

	time.Sleep(time.Second * 5)
  fmt.Println("dockerTask ->")
	fmt.Println(dockerTask, createResult)
  fmt.Println("<-dockerTask")

	fmt.Printf("stopping container %s\n", createResult.ContainerId)

	_ = stopContainer(dockerTask, createResult.ContainerId)
}
