package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/praaatik/tesseract/task"
)

// StartTaskHandler will handle the start task request from the Manager
func (a *Api) StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	a.Logger.Debug(fmt.Sprintf("StartTaskHandler reached with data - %v\n", d))

	te := task.Event{}
	err := d.Decode(&te)
	if err != nil {
		var msg = fmt.Sprintf("Error unmarshalling body: %v\n", err)
		a.Logger.Error(msg)
		w.WriteHeader(http.StatusBadRequest)
		e := ErrResponse{
			HTTPStatusCode: 400,
			Message:        msg,
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	a.Worker.AddTask(te.Task)
	a.Logger.Info(fmt.Sprintf("Added task %v\n", te.Task.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(te.Task)

	return
}

func (a *Api) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(a.Worker.GetTasks())
}

func (a *Api) StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	taskId := pathSegments[len(pathSegments)-1]

	if taskId == "" {
		a.Logger.Error("No taskID in the request\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tID, err := uuid.Parse(taskId)
	if err != nil {
		a.Logger.Error("Invalid taskID format: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskToStop, ok := a.Worker.TaskDb[tID]
	if !ok {
		a.Logger.Error("No task with ID %v found", tID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskCopy := *taskToStop
	taskCopy.State = task.Completed
	a.Worker.AddTask(taskCopy)

	a.Logger.Info("Added task %v to stop container %v\n", taskToStop.ID, taskToStop.ContainerID)
	w.WriteHeader(http.StatusNoContent)
}
