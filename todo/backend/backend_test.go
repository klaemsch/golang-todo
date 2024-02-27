package backend_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"klaemsch.io/todo/backend"
	"klaemsch.io/todo/stores"
)

func encodeJsonBody(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonData), nil
}

func TestPost(t *testing.T) {

	// create test body
	var testTodo = stores.todo{
		Id:   -100,
		Name: "Spazieren gehen!",
		Done: false,
	}

	// convert it to json
	body, err := encodeJsonBody(testTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create request for endpoint
	req, err := http.NewRequest("POST", "/todo", body)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	rr := httptest.NewRecorder()

	// make the request
	backend.PostTodo(rr, req)

	// check if server returned 200 OK as status code
	if rr.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, rr.Code)
	}

	var responseTodo stores.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatal(err)
	}

	if responseTodo.Name != testTodo.Name || responseTodo.Done != testTodo.Done {
		t.Fatal("Response data does not equal input data")
	}
}

func TestNilBodyPost(t *testing.T) {

	// create request for endpoint
	request, err := http.NewRequest("GET", "/todo", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	responseRecorder := httptest.NewRecorder()

	// do the request
	backend.PostTodo(responseRecorder, request)

	// check if server returned 400 Bad Request as status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestGet(t *testing.T) {

	// create request for endpoint
	request, err := http.NewRequest("GET", "/todo", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	responseRecorder := httptest.NewRecorder()

	// do the request
	backend.GetTodo(responseRecorder, request)

	// check if server returned 200 OK as status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, responseRecorder.Code)
	}

	var responseMap []stores.Todo
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseMap)
	if err != nil {
		t.Fatal(err)
	}

	if len(responseMap) != 0 {
		t.Fatal("Response data does not equal input data")
	}
}

func TestPostAndGetById(t *testing.T) {

	// create test body
	var testTodo = stores.Todo{
		Id:   -100,
		Name: "Spazieren gehen!",
		Done: false,
	}

	// convert it to json
	testTodoAsJson, err := encodeJsonBody(testTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create request for post endpoint
	postReq, err := http.NewRequest("POST", "/todo", testTodoAsJson)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	postRespRec := httptest.NewRecorder()

	// make the request
	backend.PostTodo(postRespRec, postReq)

	// check if server returned 200 OK as status code
	if postRespRec.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, postRespRec.Code)
	}

	var responseTodo stores.Todo
	err = json.Unmarshal(postRespRec.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatal(err)
	}

	if responseTodo.Name != testTodo.Name || responseTodo.Done != testTodo.Done {
		t.Fatal("Response data does not equal input data")
	}

	// url for the request
	url := fmt.Sprintf("/todo?id=%v", responseTodo.Id)

	// create request for get endpoint
	getReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	getRespRec := httptest.NewRecorder()

	// make the request
	backend.GetTodo(getRespRec, getReq)

	// check if server returned 200 OK as status code
	if getRespRec.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, getRespRec.Code)
	}

	err = json.Unmarshal(postRespRec.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatal(err)
	}

	if responseTodo.Name != testTodo.Name || responseTodo.Done != testTodo.Done {
		t.Fatal("Response data does not equal input data")
	}

}

func TestIdDelete(t *testing.T) {

	// create test body
	var testTodo = stores.Todo{
		Id:   -100,
		Name: "Spazieren gehen!",
		Done: false,
	}

	// convert it to json
	body, err := encodeJsonBody(testTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create a test request for the post endpoint
	postRequest, err := http.NewRequest("POST", "/todo", body)
	if err != nil {
		t.Fatal(err)
	}

	// create response recorder for the post request
	postRecorder := httptest.NewRecorder()

	// make the request
	backend.PostTodo(postRecorder, postRequest)

	// check if server returned the correct response code
	if postRecorder.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, postRecorder.Code)
	}

	var responseTodo stores.Todo
	err = json.Unmarshal(postRecorder.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create url to delete todo that was just created
	url := fmt.Sprintf("/todo?id=%v", responseTodo.Id)

	// create a test request for the delete endpoint
	deleteRequest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create response recorder for the delete request
	deleteRecorder := httptest.NewRecorder()

	// make the request
	backend.DeleteTodo(deleteRecorder, deleteRequest)

	// check if server returned the correct response code
	if deleteRecorder.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, deleteRecorder.Code)
	}

	// create request for get endpoint
	getReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	getRespRec := httptest.NewRecorder()

	// make the request
	backend.GetTodo(getRespRec, getReq)

	// check if server returned 404 Not Found as status code
	if getRespRec.Code != http.StatusNotFound {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusNotFound, getRespRec.Code)
	}
}

func TestNoIdDelete(t *testing.T) {

	// create request for endpoint
	req, err := http.NewRequest("DELETE", "/todo", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	rr := httptest.NewRecorder()

	// do the request
	backend.DeleteTodo(rr, req)

	// check if server returned 400 Bad Request as status code
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusBadRequest, rr.Code)
	}
}

func TestEmptyIdDelete(t *testing.T) {

	// create request for endpoint
	req, err := http.NewRequest("DELETE", "/todo?id=", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	rr := httptest.NewRecorder()

	// do the request
	backend.DeleteTodo(rr, req)

	// check if server returned 400 Bad Request as status code
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusBadRequest, rr.Code)
	}
}

func TestStringIdDelete(t *testing.T) {

	// create request for endpoint
	req, err := http.NewRequest("DELETE", "/todo?id=a", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	rr := httptest.NewRecorder()

	// do the request
	backend.DeleteTodo(rr, req)

	// check if server returned 400 Bad Request as status code
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusBadRequest, rr.Code)
	}
}

func TestIdNotFoundDelete(t *testing.T) {

	// create request for endpoint
	req, err := http.NewRequest("DELETE", "/todo?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	rr := httptest.NewRecorder()

	// do the request
	backend.DeleteTodo(rr, req)

	// check if server returned 400 Bad Request as status code
	if rr.Code != http.StatusNotFound {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusNotFound, rr.Code)
	}
}

func TestPostAndGetAndPut(t *testing.T) {

	// create test body
	var testTodo = stores.Todo{
		Id:   -100,
		Name: "Spazieren gehen!",
		Done: false,
	}

	// convert it to json
	testTodoAsJson, err := encodeJsonBody(testTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create request for post endpoint
	postReq, err := http.NewRequest("POST", "/todo", testTodoAsJson)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	postRespRec := httptest.NewRecorder()

	// make the request
	backend.PostTodo(postRespRec, postReq)

	// check if server returned 200 OK as status code
	if postRespRec.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, postRespRec.Code)
	}

	// create request for get endpoint
	getReq, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	getRespRec := httptest.NewRecorder()

	// make the request
	backend.GetTodo(getRespRec, getReq)

	// check if server returned 200 OK as status code
	if getRespRec.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, getRespRec.Code)
	}

	var responseMap []stores.Todo
	err = json.Unmarshal(getRespRec.Body.Bytes(), &responseMap)
	if err != nil {
		t.Fatal(err)
	}

	if len(responseMap) != 1 || responseMap[0].Name != testTodo.Name || responseMap[0].Done != testTodo.Done {
		t.Fatal("Response data does not equal input data")
	}

	// change test todo
	testTodo.Id = responseMap[0].Id
	testTodo.Done = true

	// convert it to json
	testTodoAsJson, err = encodeJsonBody(testTodo)
	if err != nil {
		t.Fatal(err)
	}

	// create request for put endpoint
	putReq, err := http.NewRequest("PUT", "/todo", testTodoAsJson)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to check responses to the test requests
	putRespRec := httptest.NewRecorder()

	// make the request
	backend.PutTodo(putRespRec, putReq)

	// check if server returned 200 OK as status code
	if putRespRec.Code != http.StatusOK {
		t.Fatalf("Got wrong status code: expected %v instead of %v", http.StatusOK, putRespRec.Code)
	}

	var responseTodo stores.Todo
	err = json.Unmarshal(putRespRec.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatal(err)
	}

	if responseTodo.Name != testTodo.Name || responseTodo.Done != testTodo.Done {
		t.Fatal("Response data does not equal input data")
	}

}
