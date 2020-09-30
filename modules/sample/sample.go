package sample

// Handler is the main structure
type Handler struct {}

// CreateSample returns a probe Handler structure
func CreateSample() *Handler {	
	return &Handler{}
}

// sayHi will return Hi value
func (H *Handler) sayHi() string {
	return "Hi"
}