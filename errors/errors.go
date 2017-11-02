package errors

type Error struct {
	HTTPStatus int
	Message    ErrorMessage
}

type ErrorMessage struct {
	Message string `json:"error"`
}

var WrongCredentials = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Wrong username or password"}}
