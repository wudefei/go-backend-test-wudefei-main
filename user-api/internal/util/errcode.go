package util

const (
	Ok       int32 = 0
	Undefine int32 = 900

	BadRequest          int32 = 400
	Invalid             int32 = 401
	Forbidden           int32 = 403
	InternalServerError int32 = 500
)

var ErrorMsg = map[int32]string{
	Ok:                  "success",
	BadRequest:          "Bad Request",
	Invalid:             "Invalid param",
	Forbidden:           "Forbidden",
	InternalServerError: "Server Internal Error",
}
