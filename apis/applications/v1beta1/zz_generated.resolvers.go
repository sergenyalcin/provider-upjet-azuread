// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Certificate.
func (mg *Certificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationObjectID")
	}
	mg.Spec.ForProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationObjectID")
	}
	mg.Spec.InitProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FederatedIdentityCredential.
func (mg *FederatedIdentityCredential) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationObjectID")
	}
	mg.Spec.ForProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationObjectID")
	}
	mg.Spec.InitProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Password.
func (mg *Password) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationObjectID")
	}
	mg.Spec.ForProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationObjectID")
	}
	mg.Spec.InitProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PreAuthorized.
func (mg *PreAuthorized) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationObjectID")
	}
	mg.Spec.ForProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizedAppID),
		Extract:      resource.ExtractParamPath("application_id", true),
		Reference:    mg.Spec.ForProvider.AuthorizedAppIDRef,
		Selector:     mg.Spec.ForProvider.AuthorizedAppIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizedAppID")
	}
	mg.Spec.ForProvider.AuthorizedAppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizedAppIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationObjectIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationObjectIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationObjectID")
	}
	mg.Spec.InitProvider.ApplicationObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizedAppID),
		Extract:      resource.ExtractParamPath("application_id", true),
		Reference:    mg.Spec.InitProvider.AuthorizedAppIDRef,
		Selector:     mg.Spec.InitProvider.AuthorizedAppIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizedAppID")
	}
	mg.Spec.InitProvider.AuthorizedAppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizedAppIDRef = rsp.ResolvedReference

	return nil
}
