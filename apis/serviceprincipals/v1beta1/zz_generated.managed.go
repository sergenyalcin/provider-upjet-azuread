// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Certificate.
func (mg *Certificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Certificate.
func (mg *Certificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Certificate.
func (mg *Certificate) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Certificate.
func (mg *Certificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Certificate.
func (mg *Certificate) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetReconciliationPolicy of this Certificate.
func (mg *Certificate) GetReconciliationPolicy() *xpv1.ReconciliationPolicy {
	return mg.Spec.ReconciliationPolicy
}

// GetWriteConnectionSecretToReference of this Certificate.
func (mg *Certificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Certificate.
func (mg *Certificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Certificate.
func (mg *Certificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Certificate.
func (mg *Certificate) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Certificate.
func (mg *Certificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Certificate.
func (mg *Certificate) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetReconciliationPolicy of this Certificate.
func (mg *Certificate) SetReconciliationPolicy(r *xpv1.ReconciliationPolicy) {
	mg.Spec.ReconciliationPolicy = r
}

// SetWriteConnectionSecretToReference of this Certificate.
func (mg *Certificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetReconciliationPolicy of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetReconciliationPolicy() *xpv1.ReconciliationPolicy {
	return mg.Spec.ReconciliationPolicy
}

// GetWriteConnectionSecretToReference of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetReconciliationPolicy of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetReconciliationPolicy(r *xpv1.ReconciliationPolicy) {
	mg.Spec.ReconciliationPolicy = r
}

// SetWriteConnectionSecretToReference of this ClaimsMappingPolicyAssignment.
func (mg *ClaimsMappingPolicyAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Password.
func (mg *Password) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Password.
func (mg *Password) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Password.
func (mg *Password) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Password.
func (mg *Password) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Password.
func (mg *Password) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetReconciliationPolicy of this Password.
func (mg *Password) GetReconciliationPolicy() *xpv1.ReconciliationPolicy {
	return mg.Spec.ReconciliationPolicy
}

// GetWriteConnectionSecretToReference of this Password.
func (mg *Password) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Password.
func (mg *Password) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Password.
func (mg *Password) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Password.
func (mg *Password) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Password.
func (mg *Password) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Password.
func (mg *Password) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetReconciliationPolicy of this Password.
func (mg *Password) SetReconciliationPolicy(r *xpv1.ReconciliationPolicy) {
	mg.Spec.ReconciliationPolicy = r
}

// SetWriteConnectionSecretToReference of this Password.
func (mg *Password) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Principal.
func (mg *Principal) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Principal.
func (mg *Principal) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Principal.
func (mg *Principal) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Principal.
func (mg *Principal) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Principal.
func (mg *Principal) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetReconciliationPolicy of this Principal.
func (mg *Principal) GetReconciliationPolicy() *xpv1.ReconciliationPolicy {
	return mg.Spec.ReconciliationPolicy
}

// GetWriteConnectionSecretToReference of this Principal.
func (mg *Principal) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Principal.
func (mg *Principal) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Principal.
func (mg *Principal) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Principal.
func (mg *Principal) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Principal.
func (mg *Principal) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Principal.
func (mg *Principal) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetReconciliationPolicy of this Principal.
func (mg *Principal) SetReconciliationPolicy(r *xpv1.ReconciliationPolicy) {
	mg.Spec.ReconciliationPolicy = r
}

// SetWriteConnectionSecretToReference of this Principal.
func (mg *Principal) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetReconciliationPolicy of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetReconciliationPolicy() *xpv1.ReconciliationPolicy {
	return mg.Spec.ReconciliationPolicy
}

// GetWriteConnectionSecretToReference of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetReconciliationPolicy of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetReconciliationPolicy(r *xpv1.ReconciliationPolicy) {
	mg.Spec.ReconciliationPolicy = r
}

// SetWriteConnectionSecretToReference of this TokenSigningCertificate.
func (mg *TokenSigningCertificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
