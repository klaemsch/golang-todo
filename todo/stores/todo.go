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
	List     *TodoList `json:"-"`
	Prev     *todo     `json:"-"`
	Next     *todo     `json:"-"`
}

type todoUpdate struct {
	todo
	UpOrDown int `json:"upOrDown"`
}

type changeOrder struct {
	Todo   *todo
	MoveUp bool
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
func UpdateTodoFromJson(r *http.Request) (*todo, *changeOrder, error) {

	// create empty todoUpdate and fill it with data from body
	todoUpdate := todoUpdate{}
	err := json.NewDecoder(r.Body).Decode(&todoUpdate)

	// create todo and fill it with data from todoUpdate
	todo := todo{
		todoUpdate.Id,
		todoUpdate.Name,
		todoUpdate.Text,
		todoUpdate.Done,
		todoUpdate.Category,
		nil,
		nil,
		nil,
	}

	if err != nil {
		return nil, nil, err
	}

	// check if the update includes a moving order
	if todoUpdate.UpOrDown == 0 {
		// return the update, no moving order
		return &todo, nil, nil
	}

	// the update includes a moving order

	moveUp := todoUpdate.UpOrDown == 1
	changeOrder := changeOrder{&todo, moveUp}
	return &todo, &changeOrder, nil
}
