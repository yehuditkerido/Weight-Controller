package metrics

import (
	"context"
	"time"
)

// ClusterMetrics holds metrics collected from a remote cluster.
type ClusterMetrics struct {
	ProviderName    string
	Healthy         bool
	LastScrapeTime  time.Time
	QueueDepth      int64
	P50LatencyMs    float64
	P99LatencyMs    float64
	VRAMUtilization float64
	ActiveRequests  int64
}

// Source defines the interface for metrics collection.
type Source interface {
	// Scrape collects metrics from the remote cluster.
	Scrape(ctx context.Context) (*ClusterMetrics, error)

	// ProviderName returns the name of the ExternalProvider.
	ProviderName() string
}
