package mails

type Error struct {
	Message string `json:"message"`
}

func ToError(error string) Error {
	return Error{
		Message: error,
	}
}
