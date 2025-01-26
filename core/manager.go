package core

import (
	"fmt"
)

var pluginRegistry = make(map[string]Plugin)

// RegisterPlugin adds a new plugin to the registry
func RegisterPlugin(plugin Plugin) {
	pluginRegistry[plugin.Name()] = plugin
}

// StartPlugin starts a specific plugin
func StartPlugin(name string) error {
	plugin, exists := pluginRegistry[name]
	if !exists {
		return fmt.Errorf("plugin %s not found", name)
	}
	return plugin.Start()
}

// StopPlugin stops a specific plugin
func StopPlugin(name string) error {
	plugin, exists := pluginRegistry[name]
	if !exists {
		return fmt.Errorf("plugin %s not found", name)
	}
	return plugin.Stop()
}

// ListPlugins returns the names of all registered plugins
func ListPlugins() []string {
	var plugins []string
	for name := range pluginRegistry {
		plugins = append(plugins, name)
	}
	return plugins
}
