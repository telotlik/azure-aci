package main

import (
	"os"
	"log"
	opencensuscli "github.com/virtual-kubelet/node-cli/opencensus"
	"github.com/virtual-kubelet/virtual-kubelet/errdefs"
	"go.opencensus.io/trace"
	"contrib.go.opencensus.io/exporter/jaeger"
)

func initJaegerAgent(c *opencensuscli.Config) (trace.Exporter, error) {
	log.Println("start register jaeger exporter")
	jOpts := jaeger.Options{
		Endpoint:      os.Getenv("JAEGER_ENDPOINT"),
		AgentEndpoint: os.Getenv("JAEGER_AGENT_ENDPOINT"),
		Username:      os.Getenv("JAEGER_USER"),
		Password:      os.Getenv("JAEGER_PASSWORD"),
		Process: jaeger.Process{
			ServiceName: c.ServiceName,
		},
	}

	if jOpts.Endpoint == "" && jOpts.AgentEndpoint == "" {
		return nil, errdefs.InvalidInput("Must specify either JAEGER_ENDPOINT or JAEGER_AGENT_ENDPOINT")
	}

	for k, v := range c.Tags {
		jOpts.Process.Tags = append(jOpts.Process.Tags, jaeger.StringTag(k, v))
	}
	log.Println("end register jaeger exporter")
	return jaeger.NewExporter(jOpts)
}
