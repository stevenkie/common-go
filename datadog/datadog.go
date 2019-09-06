package ddog

import (
	"github.com/DataDog/datadog-go/statsd"
	log "github.com/sleey/common-go/log"
)

// Datadog is datadog client wrapper
type Datadog struct {
	Client *statsd.Client
}

// Config config used for datadog package
type Config struct {
	Endpoint    string
	ServiceName string
	DefaultTags []string
}

// InitDatadog init datadog client
func InitDatadog(c Config) error {
	client, err := statsd.New(c.Endpoint)

	if err != nil {
		log.WithField("error", err).Fatal("Failed to init datadog instance")
		return err
	}

	client.Namespace = c.ServiceName + "."
	client.Tags = c.DefaultTags

	ddog = &Datadog{client}

	return nil
}
