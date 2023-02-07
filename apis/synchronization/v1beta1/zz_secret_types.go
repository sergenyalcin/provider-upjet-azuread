/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialObservation struct {
}

type CredentialParameters struct {

	// Name for this key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Value for this key-value pair.
	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type SecretObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretParameters struct {

	// +kubebuilder:validation:Optional
	Credential []CredentialParameters `json:"credential,omitempty" tf:"credential,omitempty"`

	// The object ID of the service principal for which this synchronization secret should be created
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}