package backend

import (
	"net/http"

	"klaemsch.io/todo/stores"
)

/* http Middleware for adding Cross Origin Resource Sharing Headers to the response
 * next: function / request handler that gets called after adding the headers
 * in case of OPTIONS as request method: return after the headers are set and dont call the next func
 * https://stackoverflow.com/a/64064331
 */
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request, listId string)

/* http Middleware for extracting and validating the content of the Authorization Header (token)
 * if token / listId is valid -> next function / request handler
 * if token / listId is invalid -> returns error
 */
func AUTH(next MyHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// extract the token / listId from the request header (Authorization Header)
		// returns an error if the token / listId was not found or has an incorrect format
		token, err := GetTokenFromRequest(r)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// checks if the token / listId is in the todo list store
		isValid := stores.IsValidTodoListId(token)

		if !isValid {
			// if token / listId is invalid return error
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// if token / listId is valid -> save listId in the request context
		//ctx := context.WithValue(r.Context(), "listId", token)

		// call next function with context
		//next(w, r.WithContext(ctx))
		next(w, r, token)
	}
}
