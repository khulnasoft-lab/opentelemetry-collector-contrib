// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"golang.org/x/time/rate"
)

const defaultEndpoint = ":12001"

type Config struct {
	confighttp.HTTPServerConfig `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	// Limit is a float that indicates the maximum number of messages repeated
	// through the websocket by this processor in messages per second. Defaults to 1.
	Limit rate.Limit `mapstructure:"limit"`
}

func createDefaultConfig() component.Config {
	return &Config{
		HTTPServerConfig: confighttp.HTTPServerConfig{
			Endpoint: defaultEndpoint,
		},
		Limit: 1,
	}
}
