package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TaskHandler handles HTTP requests for tasks.
type TaskHandler struct{}

// NewTaskHandler returns a configured TaskHandler.
func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

// Routes registers task routes on the chi router.
func (h *TaskHandler) Routes(r *chi.Mux) {
	r.Get("/", h.List)
	r.Post("/", h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tasks list")) // placeholder
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create task")) // placeholder
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get task")) // placeholder
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update task")) // placeholder
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete task")) // placeholder
}
