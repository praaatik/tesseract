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
}

func (a *Api) Start() {
	a.initRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%d", a.Address, a.Port), a.Router)
}

// func (a *Api) tasksHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method -> ", r.Method)
//
// 	switch r.Method {
// 	case http.MethodPost:
// 		if r.URL.Path == "/tasks/" || r.URL.Path == "/tasks" {
// 			a.StartTaskHandler(w, r)
// 			return
// 		}
// 	case http.MethodGet:
// 		if r.URL.Path == "/tasks/" || r.URL.Path == "/tasks" {
// 			a.GetTasksHandler(w, r)
// 			return
// 		}
// 	case http.MethodDelete:
// 		parts := strings.Split(r.URL.Path, "/")
// 		if len(parts) >= 3 && parts[1] == "tasks" {
// 			a.StopTaskHandler(w, r)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusMethodNotAllowed)
// }
//
// func (a *Api) initRouter() {
// 	a.Router = http.NewServeMux()
//
// 	// POST /tasks/
// 	a.Router.HandleFunc("/tasks/", a.tasksHandler)
//
// 	// DELETE /tasks/{taskID}
// 	a.Router.HandleFunc("/tasks/", a.tasksHandler)
// }
//
// func (a *Api) Start() {
// 	a.initRouter()
// 	http.ListenAndServe(fmt.Sprintf("%s:%d", a.Address, a.Port), a.Router)
// }
