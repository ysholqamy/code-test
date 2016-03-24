package event

//Result should be used for error handling. Message value should be ignored if Error != nil
type Result struct {
	Message string
	Error   error
}
