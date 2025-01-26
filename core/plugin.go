package core

import "net/http"

type Plugin interface {
	Name() string                                         // Returns the name of the plugin
	Init(config map[string]interface{}) error             // Initialize the plugin with config
	Start() error                                         // Start the plugin
	Stop() error                                          // Stop the plugin
	HandleRequest(w http.ResponseWriter, r *http.Request) // Handle API requests
}
