package task

import (
	"context"
	"io"
	"math"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/praaatik/tesseract/logger"
)

// Task is the smallest unit of work to be performed.
// This struct might be extended with more fields.
type Task struct {
	// ID is a unique identifier field for the Task.
	ID uuid.UUID

	// Human-readable name format of the Task.
	Name string

	// State is the current lifecycle state of the Task
	State State

	// Image indicates the Docker image the Task is running.
	Image string

	// Memory is useful to identify the memory the Task would require.
	Memory int

	// Disk is useful to identify the disk the Task would require.
	Disk int

	Cpu         float64
	ContainerID string

	// ExposedPorts defines the ports that the Task container will expose.
	ExposedPorts nat.PortSet

	// PortBindings maps container ports to host port.
	PortBindings map[string]string

	// RestartPolicy defines the policy which tells the system what to do when a Task fails.
	RestartPolicy string

	// StartTime is the time when the Task was started.
	StartTime time.Time

	// FinishTime is the time when the Task was completed.
	FinishTime time.Time

	Logger *logger.Logger
}

// Event represents a change in the Task state.
// Users don't interact with this, it is triggered whenever there is a change in the Task state.
// Renamed from TaskEvent to Event because of - https://go.dev/blog/package-names
type Event struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

// Config struct is used to hold the docker container configuration
type Config struct {
	// Name is used to identify the Task.
	Name string

	// AttachStdin specifies if the stdin should be attached to the Task.
	AttachStdin bool

	// AttachStdout specifies if the stdout should be attached to the Task.
	AttachStdout bool

	// AttachStderr specifies if the stderr should be attached to the Task.
	AttachStderr bool

	// ExposedPorts defines the ports that the Task container will expose.
	ExposedPorts nat.PortSet

	// Cmd specifies the Command to be executed in the container.
	Cmd []string

	// Memory is useful to identify the memory the Task would require.
	// int64 to be compatible with Docker library.
	Memory int64

	// Cpu units to be used by the Task.
	Cpu float64

	// Disk is useful to identify the disk the Task would require.
	// int64 to be compatible with Docker library.
	Disk int64

	// Env holds the environment variables to be passed to the Task.
	Env []string

	// RestartPolicy specifies the restart policy - always / unless-stopped / on-failure
	RestartPolicy string

	// Image specifies the Image the Task should run.
	Image string
}

// Docker struct is used to run a task as a Docker container
type Docker struct {
	// Client holds the Docker client used to interact with Docker API
	Client *client.Client

	// Config holds the configuration of the docker container
	Config Config
	Logger *logger.Logger
}

// DockerResult contains the result of the docker container execution
type DockerResult struct {
	// Error is used to hold error messages
	Error error

	// Action being taken, start, stop, etc
	Action string

	// ContainerId has the current ID of the container being run
	ContainerId string
	//
	// Result holds text to provide more information on the output
	Result string
}

// Run method actually runs the container
// 1. pull the Docker image from the container repository
// 2. ImagePull to pull the image
// 3. Check if ImagePull was successful
// 4. Return to standard output
// Equivalent to `docker run` command
func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	d.Logger.Info("Pulling Docker image %s", d.Config.Image)
	reader, err := d.Client.ImagePull(ctx, d.Config.Image, image.PullOptions{})

	if err != nil {
		// log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		d.Logger.Error("Failed to pull image %s: %v", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		d.Logger.Warn("Error copying image pull output: %v", err)
	}

	// Required for host configuration
	restartPolicy := container.RestartPolicy{
		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
	}

	// Required for host configuration
	resources := container.Resources{
		Memory:   d.Config.Memory,
		NanoCPUs: int64(d.Config.Cpu * math.Pow(10, 9)),
	}

	hostConfig := container.HostConfig{
		RestartPolicy:   restartPolicy,
		Resources:       resources,
		PublishAllPorts: true,
	}

	containerConfiguration := container.Config{
		Image:        d.Config.Image,
		Tty:          false,
		Env:          d.Config.Env,
		ExposedPorts: d.Config.ExposedPorts,
	}

	d.Logger.Debug("Creating container for image %s", d.Config.Image)
	resp, err := d.Client.ContainerCreate(ctx, &containerConfiguration, &hostConfig, nil, nil, d.Config.Name)
	if err != nil {
		d.Logger.Error("Failed to create container %s: %v", d.Config.Image, err)
		// log.Printf("Error creating container %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}

	d.Logger.Info("Starting container %s", resp.ID)
	err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		d.Logger.Error("Failed to start container %s: %v", resp.ID, err)
		return DockerResult{Error: err}
	}

	d.Logger.Info("Container %s started successfully", resp.ID)
	return DockerResult{
		Error:       nil,
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "success",
	}
}

func (d *Docker) Stop(id string) DockerResult {
	d.Logger.Info("Attempting to stop container %s", id)

	ctx := context.Background()

	// Check if the container exists
	_, err := d.Client.ContainerInspect(ctx, id)
	if err != nil {
		d.Logger.Warn("Container %s not found or error inspecting: %v", id, err)
		return DockerResult{Action: "stop", Result: "container not found", Error: err}
	}

	err = d.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		d.Logger.Error("Error stopping container %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	err = d.Client.ContainerRemove(ctx, id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		d.Logger.Error("Error removing container %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	d.Logger.Info("Successfully stopped and removed container %s", id)
	return DockerResult{Action: "stop", Result: "success", Error: nil}
}

func NewDocker(c *Config) *Docker {
	dc, _ := client.NewClientWithOpts(client.FromEnv)
	return &Docker{
		Client: dc,
		Config: *c,
	}
}

func NewConfig(t *Task) *Config {
	return &Config{
		Name:          t.Name,
		ExposedPorts:  t.ExposedPorts,
		Image:         t.Image,
		Cpu:           t.Cpu,
		Memory:        int64(t.Memory),
		Disk:          int64(t.Disk),
		RestartPolicy: t.RestartPolicy,
	}
}
