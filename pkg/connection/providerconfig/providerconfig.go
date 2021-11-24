package providerconfig

import (
	"context"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/fhopfensperger/provider-basic-aws/apis/v1beta1"
)

// Error strings.
const (
	errCreateOrUpdateProviderConfig = "cannot create or update connection provider config"
	errCreateOrUpdateSecret         = "cannot create or update connection secret"
)

// An Publisher publishes ConnectionDetails by submitting a Secret and ProviderConfig to a
// Kubernetes API server.
type Publisher struct {
	secret         resource.Applicator
	providerConfig resource.Applicator
	typer          runtime.ObjectTyper
}

// NewProviderConfigPublisher returns a new Publisher.
func NewProviderConfigPublisher(c client.Client, ot runtime.ObjectTyper) *Publisher {
	return &Publisher{
		secret: resource.NewApplicatorWithRetry(resource.NewAPIPatchingApplicator(c),
			resource.IsAPIErrorWrapped, nil),
		providerConfig: resource.NewApplicatorWithRetry(resource.NewAPIPatchingApplicator(c),
			resource.IsAPIErrorWrapped, nil),
		typer: ot,
	}
}

// PublishConnection publishes the supplied ConnectionDetails to a Secret in the
// same namespace as the supplied Managed resource. It is a no-op if the secret
// already exists with the supplied ConnectionDetails.
func (pcp *Publisher) PublishConnection(ctx context.Context, mg resource.Managed, c managed.ConnectionDetails) error {
	// This resource does not want to expose pcp connection secret.
	if mg.GetWriteConnectionSecretToReference() == nil {
		return nil
	}

	// Create providerconfig which references a Kubernetes secret
	pc := ConnectionProviderConfigFor(mg, resource.MustGetKind(mg, pcp.typer))
	err := pcp.providerConfig.Apply(ctx, pc, resource.MustBeControllableBy(mg.GetUID()))
	if err != nil {
		return errors.Wrap(err, errCreateOrUpdateProviderConfig)
	}

	// Create Kubernetes secret
	s := resource.ConnectionSecretFor(mg, resource.MustGetKind(mg, pcp.typer))
	s.Data = c
	// Assign the providerconfig as owner of the Kubernetes secret to avoid deletion by the Kubernetes garbage collector
	s.SetOwnerReferences(append(s.OwnerReferences, []metav1.OwnerReference{meta.AsOwner(meta.TypedReferenceTo(pc, pc.GroupVersionKind()))}...))
	err = pcp.secret.Apply(ctx, s, resource.ConnectionSecretMustBeControllableBy(mg.GetUID()))
	if err != nil {
		return errors.Wrap(err, errCreateOrUpdateSecret)
	}

	return nil
}

// UnpublishConnection is no-op since PublishConnection only creates resources
// that will be garbage collected by Kubernetes when the managed resource is
// deleted.
func (pcp *Publisher) UnpublishConnection(ctx context.Context, mg resource.Managed, c managed.ConnectionDetails) error {
	return nil
}

// ConnectionProviderConfigFor returns a connectionProvider Config
func ConnectionProviderConfigFor(o resource.ConnectionSecretOwner, kind schema.GroupVersionKind) *v1beta1.ProviderConfig {
	return &v1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       o.GetWriteConnectionSecretToReference().Namespace,
			Name:            o.GetName(),
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(o, kind))},
		},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{SecretRef: &xpv1.SecretKeySelector{
					SecretReference: xpv1.SecretReference{
						Name:      o.GetWriteConnectionSecretToReference().Name,
						Namespace: o.GetWriteConnectionSecretToReference().Namespace,
					},
					Key: "credentials",
				}},
			},
		},
	}
}
