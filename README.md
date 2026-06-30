# Weight Controller

Kubernetes controller for dynamic weight-based traffic routing across multiple clusters.

## Overview

The Weight Controller watches `ExternalModel` resources from `inference.opendatahub.io/v1alpha1` and dynamically adjusts traffic weights based on cluster metrics (latency, queue depth, health).

## Status

**Phase 1** - Controller skeleton that watches ExternalModels and logs reconciliation events.

- Phase 2: Metrics scraping from vLLM/EPP
- Phase 3: Weight calculation and patching

## Quick Start

```bash
# Build
make build

# Run locally (requires kubeconfig)
make run

# Deploy to cluster
make install

# Check logs
kubectl logs -n weight-controller-system -l app=weight-controller
```

## Development

```bash
make help      # Show all targets
make test      # Run tests
make lint      # Run linter
make build     # Full build
```

## Architecture

```
ExternalModel CR → Weight Controller → Scrape Metrics → Calculate Weights → Patch CR
                         ↑
                   Requeue every 10s
```

## License

Apache 2.0
