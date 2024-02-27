package stores

import "errors"

// array that stores all todo structs
var todoStore []todo

/* searches the todo store for todos that are associated with the given todo list id
 * returns a subarray of the todoStore with just those todos
 */
func GetTodos(listId string) []todo {
	var todosInTodoList []todo
	for _, todo := range todoStore {
		if todo.ListId == listId {
			todosInTodoList = append(todosInTodoList, todo)
		}
	}
	return todosInTodoList
}

/* searches the todo store for todo with the given id and given todo list id
 * if found -> returns a pointer to the todo with given id
 * if not found -> returns nil
 */
func GetTodoById(id int, listId string) *todo {
	for index, todo := range todoStore {
		// search for the given id
		if todo.Id == id {
			// check that the list is in the correct list
			if todo.ListId == listId {
				return &todoStore[index]
			}
			return nil
		}
	}
	// todo with given id was not found -> return nil
	return nil
}

/* appends the given todo to the todo store
 * returns the todo
 */
func AddTodo(newTodo todo) todo {
	todoStore = append(todoStore, newTodo)
	return newTodo
}

/* Updates a todo in the todo store, replaces the old todo object
 * updatedTodo:		the new todo struct that should be added
 * listId:			the id of the list the old and updated todo are asociated with
 * if successfully updated the todo -> returns update and nil-error
 * if the updated failed -> returns nil and error
 */
func UpdateTodo(updatedTodo todo, listId string) (*todo, error) {

	// try to find the old todo
	oldTodo := GetTodoById(updatedTodo.Id, listId)

	if oldTodo == nil {
		// old todo was not found -> return nil and error
		err := errors.New("todo with given id was not found")
		return nil, err
	}

	// update list id (client does not send the listId field)
	updatedTodo.ListId = oldTodo.ListId

	// update reference
	*oldTodo = updatedTodo

	return &updatedTodo, nil
}

/*
 * find a todo by id and remove it from the todo store
 * id:		id of the todo that will be deleted
 * listId:	id of the todo list the todo is ascociated with
 * returns the the deleted todo
 */
func RemoveTodo(id int, listId string) (*todo, error) {

	// iterate over todo array
	for index, curTodo := range todoStore {

		// search for the given id
		if curTodo.Id == id {

			if curTodo.ListId != listId {
				err := errors.New("todo with given id does not match the given listId")
				return nil, err
			}

			// create a new temporary array
			temp := make([]todo, 0)

			// fill it with the values of todos until index
			temp = append(temp, todoStore[:index]...)

			// fill it with the values of todos after index and assign it back to todos
			todoStore = append(temp, todoStore[index+1:]...)

			// this solution looks a bit clunky, but according to this discussion
			// (https://stackoverflow.com/a/57213476) its the best answer if the order
			// of the todos should be retained

			// return the todo that was deleted from the list
			return &curTodo, nil
		}
	}

	// todo with given id was not found
	err := errors.New("todo with given id was not found")
	return nil, err
}
