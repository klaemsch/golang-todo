package stores

import (
	"errors"
	"log"
)

/* searches the todo store for todos that are associated with the given todo list id
 * returns a subarray of the todoStore with just those todos
 * if the todo list is not found or the list is empty: an empty array is returned
 */
func GetTodos(listId string) []todo {

	// get todo list
	todoList := GetTodoListById(listId)

	// if not found or todo list is empty -> return empty array
	if todoList == nil || todoList.Start == nil {
		return make([]todo, 0)
	}

	// create todo array that will be returned
	var todosInTodoList []todo

	// iterate over the links and add every todo
	currentTodo := todoList.Start
	for currentTodo != nil {

		// for safety reasons if a falsy listid is found log the error and crash
		if currentTodo.List.Id != listId {
			log.Fatalf(
				"a link in the todo list has a different id (%v) than the list (%v)",
				currentTodo.List.Id,
				listId,
			)
		}
		// add the todo and go to the next link
		todosInTodoList = append(todosInTodoList, *currentTodo)
		currentTodo = currentTodo.Next
	}
	return todosInTodoList
}

/* searches the todo store for todo with the given id and given todo list id
 * if found -> returns a pointer to the todo with given id
 * if not found -> returns nil
 */
func GetTodoById(todoId int, listId string) *todo {

	// get todo list
	todoList := GetTodoListById(listId)

	// if not found or todo list is empty -> return nil
	if todoList == nil || todoList.Start == nil {
		return nil
	}

	// iterate over the links and search for the id
	currentTodo := todoList.Start
	for currentTodo != nil {

		// for safety reasons if a falsy listid is found log the error and crash
		if currentTodo.List.Id != listId {
			log.Fatalf(
				"a link in the todo list has a different id (%v) than the list (%v)",
				currentTodo.List.Id,
				listId,
			)
		}

		// check the id, if found return the todo
		if currentTodo.Id == todoId {
			return currentTodo
		}

		// go to the next link
		currentTodo = currentTodo.Next
	}

	// todo with given id was not found -> return nil
	return nil
}

/* appends a todo to a todo list
 * newTodo:	data of todo that will be added
 * listId: 	identifier of todo list the todo will be added to
 * returns: pointer to todo that was added
 * error: 	returns nil
 */
func AddTodo(newTodo todo, listId string) *todo {

	// search for todo list the todo will be added to
	todoList := GetTodoListById(listId)

	// if not found -> return nil
	if todoList == nil {
		return nil
	}

	// add todo list reference to the todo
	newTodo.List = todoList

	if todoList.Start == nil {
		// if the todo list has no todos yet (length = 0):
		// add new todo to the start field of the list
		todoList.Start = &newTodo
	} else {
		// if there already are todos in the list:
		// old: start -> #1 todo -> #2 todo -> ...
		// new: start -> new todo -> #1 todo -> #2 todo -> ...
		// after this, the new todo shows at the top

		// SET the next field of the new todo TO the current start todo
		newTodo.Next = todoList.Start
		// SET the prev field of the new todo TO nil because its the first
		newTodo.Prev = nil
		// SET the prev field of the current start todo TO the new todo
		todoList.Start.Prev = &newTodo
		// SET the start field of the todo list TO the address of the new todo
		todoList.Start = &newTodo
	}

	// return a pointer to the new todo
	return &newTodo
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

	// copy list reference (client does not send the listId field)
	updatedTodo.List = oldTodo.List

	// copy prev and next pointers
	updatedTodo.Prev = oldTodo.Prev
	updatedTodo.Next = oldTodo.Next

	// update reference
	*oldTodo = updatedTodo

	return &updatedTodo, nil
}

/*
 * find a todo by id and remove it from the todo store
 * todoId:	id of the todo that will be deleted
 * listId:	id of the todo list the todo is ascociated with
 * returns a pointer to the deleted todo
 */
func RemoveTodo(todoId int, listId string) (*todo, error) {

	// try to find the todo
	todoToBeDeleted := GetTodoById(todoId, listId)

	if todoToBeDeleted == nil {
		// todo was not found -> return nil and error
		err := errors.New("todo with given id was not found")
		return nil, err
	}

	if todoToBeDeleted.Next != nil {
		// there is a todo after the one that will be deleted
		// change its prev pointer to the one before the one that will be deleted
		// if its nil, then its the new first todo
		todoToBeDeleted.Next.Prev = todoToBeDeleted.Prev
	}
	if todoToBeDeleted.Prev != nil {
		// there is a todo before the one that will be deleted
		// change its next pointer to the one after the one that will be deleted
		todoToBeDeleted.Prev.Next = todoToBeDeleted.Next
	} else {
		// there is no todo before the one that will be deleted (its the first)
		// we will need to edit the todoList and its start field
		// if there is no next, the field will be nil and the list empty
		todoToBeDeleted.List.Start = todoToBeDeleted.Next
	}

	// the garbage collector should not care about this but just to be sure
	todoToBeDeleted.Prev = nil
	todoToBeDeleted.Next = nil

	// return pointer
	return todoToBeDeleted, nil
}
