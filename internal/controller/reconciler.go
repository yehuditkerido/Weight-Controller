package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	inferencev1alpha1 "github.com/opendatahub-io/ai-gateway-payload-processing/api/inference/v1alpha1"
)

// Reconciler reconciles ExternalModel resources and updates weights based on cluster metrics.
type Reconciler struct {
	client.Client
	Scheme            *runtime.Scheme
	ReconcileInterval time.Duration
}

// +kubebuilder:rbac:groups=inference.opendatahub.io,resources=externalmodels,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=inference.opendatahub.io,resources=externalmodels/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=inference.opendatahub.io,resources=externalproviders,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch

// Reconcile handles ExternalModel reconciliation.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("externalmodel", req.NamespacedName)

	var externalModel inferencev1alpha1.ExternalModel
	if err := r.Get(ctx, req.NamespacedName, &externalModel); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciling ExternalModel",
		"name", externalModel.Name,
		"namespace", externalModel.Namespace,
	)

	providerRefs := externalModel.Spec.ExternalProviderRefs
	log.Info("Found providers", "providerCount", len(providerRefs))

	// Skip single-provider models - no load balancing needed
	if len(providerRefs) < 2 {
		log.V(1).Info("Single provider, skipping weight calculation")
		return ctrl.Result{}, nil
	}

	// TODO(Phase 2): Scrape metrics from each provider
	// TODO(Phase 3): Calculate optimal weights and patch ExternalModel

	return ctrl.Result{RequeueAfter: r.ReconcileInterval}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&inferencev1alpha1.ExternalModel{}).
		Named("weight").
		Complete(r)
}
