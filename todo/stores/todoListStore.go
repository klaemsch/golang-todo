package stores

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

type todoList struct {
	Id    string
	Start *todo
}

var todoListStore []todoList

/* creates a cryptographic safe random string of 16 bytes (32 chars)
 * is used as the identifier for todo lists
 * returns random string or error
 * https://pkg.go.dev/crypto/rand#Read
 */
func randomString() (string, error) {

	// create empty byte array of size 16
	b := make([]byte, 16)

	// fill byte array with random bytes
	_, err := rand.Read(b)

	if err != nil {
		log.Println("error:", err)
		return "", err
	}

	// return encoded hex string from byte array
	return hex.EncodeToString(b), nil
}

/* creates a new todo list and adds it to the todo list db
 * is used as the identifier for todo lists
 * returns id of todo list or error
 */
func NewTodoList() (string, error) {

	// create new todo list identifier
	todoListId, err := randomString()

	if err != nil {
		return "", err
	}

	// create new todo list struct with generated id and empty start field
	// (will be added when the first todo is added to th list)
	todoList := todoList{
		Id:    todoListId,
		Start: nil,
	}

	// append new todo list to todo list db
	todoListStore = append(todoListStore, todoList)

	// return id of new todo list
	return todoListId, nil
}

/* checks the todo list database and confirms/denies
 * that the given todo list exists by comparing the identifiers
 * if valid -> returns true
 * if invalid -> returns false
 */
func IsValidTodoListId(listId string) bool {
	for _, todoList := range todoListStore {
		if listId == todoList.Id {
			return true
		}
	}
	return false
}

/* searches the todo list store for todo list with the given id
 * if found -> returns a pointer to the todo list with given id
 * if not found -> returns nil
 */
func GetTodoListById(listId string) *todoList {
	for index, todoList := range todoListStore {
		// search for the given id
		if todoList.Id == listId {
			return &todoListStore[index]
		}
	}
	// todo list with given id was not found -> return nil
	return nil
}
