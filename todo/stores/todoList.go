package stores

import "errors"

type TodoList struct {
	Id    string
	Start *todo
}

// uses the todo list links to create a slice of all todos
func (todoList *TodoList) GetTodos() []todo {

	// create todo array that will be returned
	var todosInTodoList []todo

	// iterate over the links and add every todo
	currentTodo := todoList.Start
	for currentTodo != nil {
		// add the todo and go to the next link
		todosInTodoList = append(todosInTodoList, *currentTodo)
		currentTodo = currentTodo.Next
	}
	return todosInTodoList
}

/* searches the todo list for todo with the given id
 * if found -> returns a pointer to the todo with given id
 * if not found -> returns nil
 */
func (todoList *TodoList) GetTodoById(todoId int) *todo {

	// iterate over the links and search for the id
	currentTodo := todoList.Start
	for currentTodo != nil {

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
 * returns: pointer to todo that was added
 * error: 	returns nil
 */
func (todoList *TodoList) AddTodo(newTodo todo) *todo {

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

/* Updates a todo in the todo list, replaces the old todo object
 * updatedTodo:		the new todo struct that should be added
 * if successfully updated the todo -> returns update and nil-error
 * if the updated failed -> returns nil and error
 */
func (todoList *TodoList) UpdateTodo(updatedTodo todo) (*todo, error) {

	// try to find the old todo
	oldTodo := todoList.GetTodoById(updatedTodo.Id)

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
 * find a todo by id and remove it from the todo list
 * todoId:	id of the todo that will be deleted
 * returns a pointer to the deleted todo
 */
func (todoList *TodoList) RemoveTodo(todoId int) (*todo, error) {

	// try to find the todo
	todoToBeDeleted := todoList.GetTodoById(todoId)

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

/*
 * find a todo by id and moves it up or down in the todo list order
 * todoId:	id of the todo that will be moved
 * moveUp:	if true moves the todo one position up, if false one position down
 * returns:	a pointer to the moved todo
 */
func (todoList *TodoList) MoveTodo(todoId int, moveUp bool) (*todo, error) {

	// try to find the todo
	todoToBeMoved := todoList.GetTodoById(todoId)

	if todoToBeMoved == nil {
		// todo was not found -> return nil and error
		err := errors.New("todo with given id was not found")
		return nil, err
	}

	if moveUp {
		return todoList.MoveTodoUp(todoToBeMoved)
	} else {
		return todoList.MoveTodoDown(todoToBeMoved)
	}
}

/*
 * gets a todo and moves it up in the todo list order
 * todoToBeMoved:	id of the todo that will be moved
 * returns:			a pointer to the moved todo
 */
func (todoList *TodoList) MoveTodoUp(todoToBeMoved *todo) (*todo, error) {

	// if todo is the top most todo in the list, do nothing
	if todoToBeMoved.Prev == nil {
		return todoToBeMoved, nil
	}

	// get the previous todo
	prevTodo := todoToBeMoved.Prev

	// if the previous todo of prevTodo is nil, prevTodo was the start todo in the list
	if prevTodo.Prev == nil {
		// set todoToBeMoved as start of the list
		todoList.Start = todoToBeMoved
	} else {
		// previous todo was not the first so get its previous and set todoToBeMoved as next
		prevTodo.Prev.Next = todoToBeMoved
	}

	// swap todoToBeMoved with its previous todo
	prevTodo.Next, todoToBeMoved.Next = todoToBeMoved.Next, prevTodo
	todoToBeMoved.Prev, prevTodo.Prev = prevTodo.Prev, todoToBeMoved

	// return pointer
	return todoToBeMoved, nil
}

/*
 * gets a todo and moves it down in the todo list order
 * todoToBeMoved:	id of the todo that will be moved
 * returns:			a pointer to the moved todo
 */
func (todoList *TodoList) MoveTodoDown(todoToBeMoved *todo) (*todo, error) {

	// if todo is the bottom most todo in the list, do nothing
	if todoToBeMoved.Next == nil {
		return todoToBeMoved, nil
	}

	// get the next todo
	nextTodo := todoToBeMoved.Next

	// if prev of todoToBeMoved is nil, todoToBeMoved was the start todo in the list
	if todoToBeMoved.Prev == nil {
		// set nextTodo as start of the list
		todoList.Start = nextTodo
	} else {
		// todoToBeMoved was not the first so get its prev and set nextTodo as next
		todoToBeMoved.Prev.Next = nextTodo
	}

	// swap todoToBeMoved with its next todo
	todoToBeMoved.Next, nextTodo.Next = nextTodo.Next, todoToBeMoved
	todoToBeMoved.Prev, nextTodo.Prev = nextTodo, todoToBeMoved.Prev

	// return pointer
	return todoToBeMoved, nil
}
