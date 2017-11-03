package errors

type Error struct {
	HTTPStatus int
	Message    ErrorMessage
}

type ErrorMessage struct {
	Message string `json:"error"`
}

var WrongCredentials = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Wrong username or password"}}
var InvalidParameters = Error{HTTPStatus: 422, Message: ErrorMessage{Message: "Invalid Parameters"}}
var SearchForShop = Error{HTTPStatus: 400, Message: ErrorMessage{Message: "You need to search for near shops first"}}
