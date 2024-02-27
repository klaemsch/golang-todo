package main

import (
	"log"
	"net/http"

	"klaemsch.io/todo/backend"
)

/*
 * found this post / tutorial: https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
 * decided it was too complicated for a first project
 * but could be interesting for another time
 * there is also a better way of writing middleware than to chain the function calls,
 * but for now this way is sufficient
 */

// https://pkg.go.dev/net/http#hdr-Patterns
func main() {
	http.HandleFunc("OPTIONS /api/todo", backend.CORS(nil))
	http.HandleFunc("GET /api/todo", backend.CORS(backend.AUTH(backend.GetTodo)))
	http.HandleFunc("POST /api/todo", backend.CORS(backend.AUTH(backend.PostTodo)))
	http.HandleFunc("PUT /api/todo", backend.CORS(backend.AUTH(backend.PutTodo)))
	http.HandleFunc("DELETE /api/todo", backend.CORS(backend.AUTH(backend.DeleteTodo)))
	http.HandleFunc("OPTIONS /api/list", backend.CORS(nil))
	http.HandleFunc("GET /api/list", backend.CORS(backend.NewTodoList))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
