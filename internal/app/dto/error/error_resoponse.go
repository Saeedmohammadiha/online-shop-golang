package dto

type ErrorData struct {
	Error   error
	Message string
	Status  int
}
type Error struct {
	Data ErrorData
}
