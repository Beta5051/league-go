package lcu

type Error struct {
	ErrorCode  string `json:"errorCode"`
	HttpStatus uint   `json:"httpStatus"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
