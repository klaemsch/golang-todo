package backend

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/* Uses a request object to retrieve the id parameter of the request url
 * the id parameter is used to identify todos
 * r: request
 * if id found: returns retrieved id and nil-error
 * if id not found: returns -1 as id and error
 */
func getIdFromUrl(r *http.Request) (int, error) {

	// use url package to get params from query/url
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err = errors.New("error while parsing url query")
		return -1, err
	}

	// try to extract the id parameter (as a string)
	idString := params.Get("id")

	// no value for id in the url parameters -> create error and return
	if idString == "" {
		err = errors.New("no ID found in the URL")
		return -1, err
	}

	// convert idString into an integer
	idValue, err := strconv.Atoi(idString)

	// atoi throws error if value for id was not a number -> create error and return
	if err != nil {
		err = errors.New("value for id was not a number")
		return -1, err
	}

	return idValue, nil
}

/* In this app, the listId is used as an authorization token
 * this token has to be send as an Autorization Bearer Header with the request to access the todo data
 *
 * Extracts the Authorization Header from the request and validates the format of the token / listId
 * IMPORTANT: does only validate that the token has the correct format, not if the token is known
 * r: request
 * if token found and format-valid:			returns token and nil-error
 * if token not found or format-invalid:	returns empty string and error
 */
func GetTokenFromRequest(r *http.Request) (string, error) {

	// get authorization header from request and split the 'Bearer '-Part (whitespace!)
	authHeader := r.Header.Get("Authorization")
	splittedAuthHeader := strings.Split(authHeader, "Bearer ")

	// the splitted version should have length 2, if not return error
	if len(splittedAuthHeader) != 2 {
		err := errors.New("incorrect authorization method")
		return "", err
	}

	// the token should contain 32 chars, if not return error
	token := splittedAuthHeader[1]
	if len(token) != 32 {
		err := errors.New("incorrect authorization method")
		return "", err
	}

	return token, nil
}
