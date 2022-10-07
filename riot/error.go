package riot

type Error struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

func (e Error) Error() string {
	return e.Message
}
