// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package diskscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"

import (
	"context"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper/internal/metadata"
)

// This file implements Factory for Disk scraper.

const (
	// TypeStr the value of "type" key in configuration.
	TypeStr = "disk"
)

// Factory is the Factory for scraper.
type Factory struct{}

// CreateDefaultConfig creates the default configuration for the Scraper.
func (f *Factory) CreateDefaultConfig() internal.Config {
	return &Config{
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

// CreateMetricsScraper creates a scraper based on provided config.
func (f *Factory) CreateMetricsScraper(
	ctx context.Context,
	settings receiver.Settings,
	config internal.Config,
) (scraper.Metrics, error) {
	cfg := config.(*Config)
	s, err := newDiskScraper(ctx, settings, cfg)
	if err != nil {
		return nil, err
	}

	return scraper.NewMetrics(s.scrape,
		scraper.WithStart(s.start),
	)
}
