package dto

type ErrorData struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type Error struct {
	Data ErrorData `json:"data"`
}
