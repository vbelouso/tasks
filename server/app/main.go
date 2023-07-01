package main

import (
	"log"
	"net/http"

	"github.com/vbelouso/tasks/server/gen/tasks"
)

func main() {
	taskService := tasks.NewTasksApiService()
	handler := tasks.NewRouter(tasks.NewTasksApiController(taskService))
	address := ":8080"
	log.Fatal(http.ListenAndServe(address, handler))
}
