package ddog

import (
	"github.com/sirupsen/logrus"
	log "github.com/sleey/common-go/log"
)

// Histogram send histogram data to datadog
func Histogram(name string, requestTime float64, tags []string, rate float64) {
	if ddog == nil {
		log.Error("datadog client is not initialized")
		return
	}

	err := ddog.Client.Histogram(name, requestTime, tags, rate)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
			"name":  name,
		}).Error("Failed to send histogram data to datadog")
	}
}

// Gauge send gauge data to datadog
func Gauge(name string, value float64, tags []string, rate float64) {
	if ddog == nil {
		log.Error("datadog client is not initialized")
		return
	}

	err := ddog.Client.Gauge(name, value, tags, rate)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
			"name":  name,
		}).Error("Failed to send gauge data to datadog")
	}
}

// Count send count data to datadog
func Count(name string, value int64, tags []string, rate float64) {
	if ddog == nil {
		log.Error("datadog client is not initialized")
		return
	}

	err := ddog.Client.Count(name, value, tags, rate)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
			"name":  name,
		}).Error("Failed to send count data to datadog")
	}
}
