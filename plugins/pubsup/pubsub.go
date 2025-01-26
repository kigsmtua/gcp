package pubsub

import (
	"fmt"
	"gcp-localstack/core"
	"net/http"
	"sync"
)

type PubSubPlugin struct {
	topics map[string][]string
	lock   sync.Mutex
}

func NewPubSubPlugin() core.Plugin {
	return &PubSubPlugin{
		topics: make(map[string][]string),
	}
}

func (p *PubSubPlugin) Name() string {
	return "pubsub"
}

func (p *PubSubPlugin) Init(config map[string]interface{}) error {
	fmt.Println("Initializing Pub/Sub plugin...")
	return nil
}

func (p *PubSubPlugin) Start() error {
	fmt.Println("Starting Pub/Sub plugin...")
	http.HandleFunc("/pubsub/topics", p.HandleRequest)
	go http.ListenAndServe(":8085", nil)
	return nil
}

func (p *PubSubPlugin) Stop() error {
	fmt.Println("Stopping Pub/Sub plugin...")
	return nil
}

func (p *PubSubPlugin) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		topic := r.URL.Query().Get("topic")
		if topic == "" {
			http.Error(w, "Missing topic parameter", http.StatusBadRequest)
			return
		}
		p.lock.Lock()
		defer p.lock.Unlock()
		p.topics[topic] = []string{}
		fmt.Fprintf(w, "Topic %s created successfully", topic)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
