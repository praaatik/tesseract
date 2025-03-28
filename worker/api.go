package worker

import (
	"fmt"
	"net/http"

	"github.com/praaatik/tesseract/logger"
)

type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

type Api struct {
	Address string
	Port    int
	Worker  *Worker
	Logger  *logger.Logger
	Router  *http.ServeMux
}

func (a *Api) initRouter() {
	a.Router = http.NewServeMux()

	// Task creation
	a.Router.HandleFunc("POST /tasks", a.StartTaskHandler)

	// Getting new tasks
	a.Router.HandleFunc("GET /tasks", a.GetTasksHandler)

	// Stopping tasks (deleting)
	a.Router.HandleFunc("DELETE /tasks/{taskID}", a.StopTaskHandler)

	// Get the statistics
	a.Router.HandleFunc("/stats", a.StatsHandler)
}

func (a *Api) Start() {
	a.initRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%d", a.Address, a.Port), a.Router)
}
