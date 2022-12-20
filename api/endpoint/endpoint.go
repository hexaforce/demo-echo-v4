package handler

// Endpoint
type Endpoint struct {
}

// NewEndpoint
func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

// Message
type Message struct {
	Message string `json:"message" example:"message"`
}
