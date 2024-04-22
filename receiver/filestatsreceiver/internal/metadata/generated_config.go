// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for filestats metrics.
type MetricsConfig struct {
	FileAtime MetricConfig `mapstructure:"file.atime"`
	FileCount MetricConfig `mapstructure:"file.count"`
	FileCtime MetricConfig `mapstructure:"file.ctime"`
	FileMtime MetricConfig `mapstructure:"file.mtime"`
	FileSize  MetricConfig `mapstructure:"file.size"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		FileAtime: MetricConfig{
			Enabled: false,
		},
		FileCount: MetricConfig{
			Enabled: false,
		},
		FileCtime: MetricConfig{
			Enabled: false,
		},
		FileMtime: MetricConfig{
			Enabled: true,
		},
		FileSize: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool            `mapstructure:"enabled"`
	Include []filter.Config `mapstructure:"include"`
	Exclude []filter.Config `mapstructure:"exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for filestats resource attributes.
type ResourceAttributesConfig struct {
	FileName ResourceAttributeConfig `mapstructure:"file.name"`
	FilePath ResourceAttributeConfig `mapstructure:"file.path"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		FileName: ResourceAttributeConfig{
			Enabled: true,
		},
		FilePath: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for filestats metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
