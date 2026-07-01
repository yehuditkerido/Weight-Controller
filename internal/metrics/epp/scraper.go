package epp

import (
	"context"
	"time"

	"github.com/yehuditkerido/Weight-Controller/internal/metrics"
)

const DefaultTimeout = 5 * time.Second

var _ metrics.Source = (*Scraper)(nil)

// Scraper collects metrics from EPP (Endpoint Picker Plugin) instances.
type Scraper struct {
	providerName string
	endpoint     string
	timeout      time.Duration
}

// Option configures a Scraper.
type Option func(*Scraper)

// WithTimeout sets the HTTP timeout for scraping.
func WithTimeout(timeout time.Duration) Option {
	return func(s *Scraper) {
		if timeout > 0 {
			s.timeout = timeout
		}
	}
}

// NewScraper creates an EPP metrics scraper.
func NewScraper(providerName, endpoint string, opts ...Option) *Scraper {
	s := &Scraper{
		providerName: providerName,
		endpoint:     endpoint,
		timeout:      DefaultTimeout,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// ProviderName returns the provider name.
func (s *Scraper) ProviderName() string {
	return s.providerName
}

// Scrape collects metrics from the EPP Prometheus endpoint.
// TODO(Phase 2): Implement actual Prometheus scraping.
func (s *Scraper) Scrape(ctx context.Context) (*metrics.ClusterMetrics, error) {
	_ = ctx // Will be used for HTTP request cancellation

	return &metrics.ClusterMetrics{
		ProviderName:   s.providerName,
		Healthy:        true,
		LastScrapeTime: time.Now(),
	}, nil
}
