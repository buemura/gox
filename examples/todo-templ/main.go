package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/buemura/gox/examples/todo-templ/views"
)

var (
	mu     sync.Mutex
	todos  []views.Todo
	nextID = 1
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/toggle", handleToggle)
	http.HandleFunc("/delete", handleDelete)

	fmt.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	mu.Lock()
	snapshot := make([]views.Todo, len(todos))
	copy(snapshot, todos)
	mu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.TodoList(snapshot).Render(r.Context(), w)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if text != "" {
		mu.Lock()
		todos = append(todos, views.Todo{ID: nextID, Text: text})
		nextID++
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleToggle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	mu.Lock()
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = !todos[i].Done
			break
		}
	}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	mu.Lock()
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
