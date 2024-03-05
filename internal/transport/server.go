package transport

type Server interface {
	// Start starts server and return error
	Start() error
	// Shutdown shutdowns server gracefully
	Shutdown() error
}
