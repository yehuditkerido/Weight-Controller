package health

import "context"

// State represents the health state of a cluster.
type State string

const (
	StateHealthy   State = "Healthy"
	StateUnhealthy State = "Unhealthy"
	StateUnknown   State = "Unknown"
)

// Source provides health status for clusters.
// TODO: Implement once the health monitoring task is complete.
type Source interface {
	// GetState returns the health state for a provider.
	GetState(ctx context.Context, providerName string) (State, error)

	// IsHealthy returns true if the cluster is healthy.
	IsHealthy(ctx context.Context, providerName string) (bool, error)
}

// PlaceholderSource is a stub that assumes all clusters are healthy.
type PlaceholderSource struct{}

// NewPlaceholderSource creates a placeholder health source.
func NewPlaceholderSource() Source {
	return &PlaceholderSource{}
}

// GetState returns healthy for all providers (placeholder).
func (p *PlaceholderSource) GetState(ctx context.Context, providerName string) (State, error) {
	return StateHealthy, nil
}

// IsHealthy returns true for all providers (placeholder).
func (p *PlaceholderSource) IsHealthy(ctx context.Context, providerName string) (bool, error) {
	return true, nil
}
