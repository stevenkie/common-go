package tracer

import (
	"context"
	"runtime"
	"strings"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

// Config used to set jaeger config
type Config struct {
	ServiceName    string
	SamplerConfig  *jaegerConfig.SamplerConfig
	ReporterConfig *jaegerConfig.ReporterConfig
}

// Initialize init tracer
func Initialize(c Config) (err error) {
	samplerConfig := &jaegerConfig.SamplerConfig{
		Type:  "probabilistic",
		Param: 0.25,
	}

	reporterConfig := &jaegerConfig.ReporterConfig{
		LogSpans:            true,
		BufferFlushInterval: 1 * time.Second,
		LocalAgentHostPort:  "localhost:6831",
	}

	if c.SamplerConfig != nil {
		samplerConfig = c.SamplerConfig
	}

	if c.ReporterConfig != nil {
		reporterConfig = c.ReporterConfig
	}

	cfg := jaegerConfig.Configuration{
		Sampler:  samplerConfig,
		Reporter: reporterConfig,
	}

	jLogger := jaegerlog.NullLogger

	_, err = cfg.InitGlobalTracer(
		c.ServiceName,
		jaegerConfig.Logger(jLogger),
	)

	if err != nil {
		return
	}

	return
}

// StartSpanFromContext create span with function name as default span name
func StartSpanFromContext(ctx context.Context) (opentracing.Span, context.Context) {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	operationName := functionName[strings.LastIndex(functionName, "/")+1:]
	return opentracing.StartSpanFromContext(ctx, operationName)
}

// StartSpanWithName to give span custom name
func StartSpanWithName(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName)
}
