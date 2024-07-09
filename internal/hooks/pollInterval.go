package hooks

import (
	"context"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"math/rand"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/upbound/provider-azuread/apis/v1beta1"
)

// opts = append(opts, hooks.WithCustomPollInterval(mgr.GetClient(), o.Logger, o.PollJitter))

func WithCustomPollInterval(client client.Client, log logging.Logger, jitter time.Duration) managed.ReconcilerOption {
	return managed.WithPollIntervalHook(func(mg resource.Managed, pollInterval time.Duration) time.Duration {
		reconciliationPolicy := mg.GetReconciliationPolicy()
		if reconciliationPolicy != nil {
			duration, err := time.ParseDuration(reconciliationPolicy.PollInterval)
			if err != nil {
				log.Debug("cannot parse custom poll interval", "error", err)
				return applyJitter(pollInterval, jitter)
			}
			return applyJitter(duration, jitter)
		}

		configRef := mg.GetProviderConfigReference()
		pc := &v1beta1.ProviderConfig{}
		err := client.Get(context.TODO(), types.NamespacedName{Name: configRef.Name}, pc)
		if err != nil {
			return applyJitter(pollInterval, jitter)
		}

		gvk := mg.GetObjectKind().GroupVersionKind()
		for _, rp := range pc.Spec.ReconciliationPolicies {
			if matchesPolicy(rp.Target, gvk, mg.GetName()) {
				return applyJitter(rp.PollInterval.Duration, jitter)
			}
		}

		return applyJitter(pollInterval, jitter)
	})
}

func applyJitter(duration, jitter time.Duration) time.Duration {
	return duration + time.Duration((rand.Float64()-0.5)*2*float64(jitter)) //nolint:gosec // No need for secure randomness.
}

func matchesPolicy(target v1beta1.GroupKindName, gvk schema.GroupVersionKind, name string) bool {
	if target.Name != nil && target.Kind != nil {
		if target.Group == gvk.Group && *target.Kind == gvk.Kind && *target.Name == name {
			return true
		}
	}
	if target.Kind != nil {
		if target.Group == gvk.Group && *target.Kind == gvk.Kind {
			return true
		}
	}
	return target.Group == gvk.Group
}
