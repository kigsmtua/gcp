package main

import (
	"gcp-localstack/cmd"
	"gcp-localstack/core"
	"gcp-localstack/plugins/pubsub"
)

func main() {
	core.RegisterPlugin(pubsub.NewPubSubPlugin())
	cmd.Execute()
}
