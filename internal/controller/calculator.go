package controller

import (
	"github.com/yehuditkerido/Weight-Controller/internal/metrics"
)

const (
	DefaultHealthWeight     = 0.4
	DefaultLatencyWeight    = 0.3
	DefaultQueueDepthWeight = 0.3
)

// ProviderWeight represents the calculated weight for a provider.
type ProviderWeight struct {
	ProviderName string
	Weight       int
}

// Calculator computes optimal weights based on cluster metrics.
type Calculator struct {
	HealthWeight     float64
	LatencyWeight    float64
	QueueDepthWeight float64
}

// CalculatorOption configures a Calculator.
type CalculatorOption func(*Calculator)

// WithHealthWeight sets the weight for health score in calculation.
func WithHealthWeight(w float64) CalculatorOption {
	return func(c *Calculator) {
		if w >= 0 && w <= 1 {
			c.HealthWeight = w
		}
	}
}

// WithLatencyWeight sets the weight for latency in calculation.
func WithLatencyWeight(w float64) CalculatorOption {
	return func(c *Calculator) {
		if w >= 0 && w <= 1 {
			c.LatencyWeight = w
		}
	}
}

// WithQueueDepthWeight sets the weight for queue depth in calculation.
func WithQueueDepthWeight(w float64) CalculatorOption {
	return func(c *Calculator) {
		if w >= 0 && w <= 1 {
			c.QueueDepthWeight = w
		}
	}
}

// NewCalculator creates a Calculator with the given options.
func NewCalculator(opts ...CalculatorOption) *Calculator {
	c := &Calculator{
		HealthWeight:     DefaultHealthWeight,
		LatencyWeight:    DefaultLatencyWeight,
		QueueDepthWeight: DefaultQueueDepthWeight,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// CalculateWeights computes weights for each provider based on their metrics.
// TODO(Phase 3): Implement actual weight calculation algorithm.
func (c *Calculator) CalculateWeights(clusterMetrics map[string]metrics.ClusterMetrics) []ProviderWeight {
	if len(clusterMetrics) == 0 {
		return nil
	}

	results := make([]ProviderWeight, 0, len(clusterMetrics))
	for name := range clusterMetrics {
		results = append(results, ProviderWeight{
			ProviderName: name,
			Weight:       1, // Default equal weight until algorithm is implemented
		})
	}
	return results
}
