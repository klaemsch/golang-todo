package stores

import (
	"encoding/json"
	"net/http"
)

// current id of last todo
var index int = 0

// structure of a todo
// not public, use constructor functions below
type todo struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Text     string    `json:"text"`
	Done     bool      `json:"done"`
	Category []string  `json:"category"`
	List     *todoList `json:"-"`
	Prev     *todo     `json:"-"`
	Next     *todo     `json:"-"`
}

/* Create a new Todo
 * r:		request with json body
 * returns: pointer to the temporary todo or error
 */
func NewTodoFromJson(r *http.Request) (*todo, error) {

	// create empty todo and fill it with data from request body
	newTodo := todo{}
	err := json.NewDecoder(r.Body).Decode(&newTodo)

	if err != nil {
		return nil, err
	}

	// insert id and update index
	newTodo.Id = index
	index++

	return &newTodo, err
}

/* Convert request body to a new, temporary Todo, that does not get inserted to db
 * r:		request with json body
 * returns: pointer to the temporary todo or error
 */
func JsonToTodo(r *http.Request) (*todo, error) {

	// create empty todo and fill it with data from body
	todo := todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, err
}
