package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"klaemsch.io/todo/stores"
)

/*
 * Creates a new todo list, adds the list to the todo list store and returns the listId
 */
func NewTodoList(w http.ResponseWriter, r *http.Request) {

	// create new todo list
	listId, err := stores.NewTodoList()
	if err != nil {
		log.Printf("Error while creating new todo list: %v", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// return list id of new todo list as simple plain text response
	fmt.Fprintln(w, listId)
}

/*
 * Returns data of
 * a) all todos if no id is given
 * b) only the todo with given id (id as url parameter) -> calls GetTodoById
 */
func GetTodo(w http.ResponseWriter, r *http.Request, todoList *stores.TodoList) {

	// regex pattern to find ?id=
	pattern := `\?id=([\w-]*)`

	// compile regex
	re := regexp.MustCompile(pattern)

	// instead of regex we could use the helper function getIdFromUrl to find out
	// wether or not a correct id was given
	// problem with this approach is, that invalid id-requests like ?id=a
	// would return a response with 200 OK and all todos as body
	// this would be a problem for debugging

	// run regex and try to find the pattern in the url path
	if re.FindStringIndex(r.URL.String()) == nil {
		// if the pattern was not found: return all todos
		json.NewEncoder(w).Encode(todoList.GetTodos())
		log.Println("GET /todo (200 OK)")
	} else {
		// if the pattern was found: pass to "byId"-function
		GetTodoById(w, r, todoList)
	}
}

/*
 * Returns todo with id given in url parameter
 */
func GetTodoById(w http.ResponseWriter, r *http.Request, todoList *stores.TodoList) {

	// get id from url
	idValue, err := getIdFromUrl(r)

	// if an error occurs while extracting the id from the url -> 400 Bad Request
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get todo with given id from db
	todo := todoList.GetTodoById(idValue)

	if todo == nil {
		log.Printf("todo with id %v not found", idValue)
		http.Error(w, "todo with given id not found", http.StatusNotFound)
	} else {
		// encode to json and send back
		json.NewEncoder(w).Encode(todo)
		log.Printf("GET /todo?id=%v (200 OK)", idValue)
	}
}

/*
 * Uses the request body to create a new todo in the todoStore
 * returns the posted todo
 */
func PostTodo(w http.ResponseWriter, r *http.Request, todoList *stores.TodoList) {

	// check if body is empty -> send 400 Bad Request back
	if r.Body == nil {
		log.Println("Request body is nil")
		http.Error(w, "Request body is nil", http.StatusBadRequest)
		return
	}

	// decode json from request body
	todo, err := stores.NewTodoFromJson(r)

	if err != nil {
		log.Printf("Error while decoding body: %v", err)
		http.Error(w, "Body is corrupted", http.StatusBadRequest)
		return
	}

	// add body=todo to the todo-database
	newTodo := todoList.AddTodo(*todo)

	// send posted todo back
	json.NewEncoder(w).Encode(newTodo)
	log.Println("POST /todo (200 OK)")
}

/*
 * Uses the request body to update the data of a todo in the todostore
 * returns the updated todo
 */
func PutTodo(w http.ResponseWriter, r *http.Request, todoList *stores.TodoList) {

	// check if body is empty -> send 400 Bad Request back
	if r.Body == nil {
		log.Println("Request body is nil")
		http.Error(w, "Request body is nil", http.StatusBadRequest)
		return
	}

	// decode json from request body
	updatedTodo, changeOrder, err := stores.UpdateTodoFromJson(r)

	if err != nil {
		log.Printf("Error while decoding body: %v", err)
		http.Error(w, "Body is corrupted", http.StatusBadRequest)
		return
	}

	fmt.Print(updatedTodo)

	// update todo in the database
	updatedTodo, err = todoList.UpdateTodo(*updatedTodo)

	if err != nil {
		log.Printf("Todo with id %v not found, could not be updated", (*updatedTodo).Id)
		http.Error(w, "Todo with given id not found", http.StatusBadRequest)
	}

	if changeOrder != nil {
		updatedTodo, err = todoList.MoveTodo(updatedTodo.Id, changeOrder.MoveUp)
	}

	if err != nil {
		log.Printf("Error while moving todo: %v", err)
		http.Error(w, "Error while moving todo", http.StatusInternalServerError)
		return
	}

	// send updated todo back
	json.NewEncoder(w).Encode(*updatedTodo)
	log.Printf("PUT /todo?id=%v (200 OK)", (*updatedTodo).Id)
}

/*
 * Uses the given id to delete the corresponding todo in the todo store
 * returns the todo that was deleted
 */
func DeleteTodo(w http.ResponseWriter, r *http.Request, todoList *stores.TodoList) {

	// get id from url
	idValue, err := getIdFromUrl(r)

	// if an error occurs while extracting the id from the url -> 400 Bad Request
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// remove todo with given id
	removedTodo, err := todoList.RemoveTodo(idValue)

	if err != nil {
		log.Printf("Todo with id %v not found, could not be deleted", idValue)
		http.Error(w, "Todo with given id not found", http.StatusNotFound)
	} else {
		// send success message back
		json.NewEncoder(w).Encode(*removedTodo)
		log.Printf("DELETE /todo?id=%v (200 OK)", idValue)
	}
}
