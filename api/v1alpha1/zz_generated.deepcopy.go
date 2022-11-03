//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Classifier) DeepCopyInto(out *Classifier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Classifier.
func (in *Classifier) DeepCopy() *Classifier {
	if in == nil {
		return nil
	}
	out := new(Classifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Classifier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierLabel) DeepCopyInto(out *ClassifierLabel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierLabel.
func (in *ClassifierLabel) DeepCopy() *ClassifierLabel {
	if in == nil {
		return nil
	}
	out := new(ClassifierLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierList) DeepCopyInto(out *ClassifierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Classifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierList.
func (in *ClassifierList) DeepCopy() *ClassifierList {
	if in == nil {
		return nil
	}
	out := new(ClassifierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReport) DeepCopyInto(out *ClassifierReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReport.
func (in *ClassifierReport) DeepCopy() *ClassifierReport {
	if in == nil {
		return nil
	}
	out := new(ClassifierReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportList) DeepCopyInto(out *ClassifierReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClassifierReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportList.
func (in *ClassifierReportList) DeepCopy() *ClassifierReportList {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportSpec) DeepCopyInto(out *ClassifierReportSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportSpec.
func (in *ClassifierReportSpec) DeepCopy() *ClassifierReportSpec {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportStatus) DeepCopyInto(out *ClassifierReportStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportStatus.
func (in *ClassifierReportStatus) DeepCopy() *ClassifierReportStatus {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierSpec) DeepCopyInto(out *ClassifierSpec) {
	*out = *in
	if in.DeployedResources != nil {
		in, out := &in.DeployedResources, &out.DeployedResources
		*out = make([]DeployedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesVersionConstraints != nil {
		in, out := &in.KubernetesVersionConstraints, &out.KubernetesVersionConstraints
		*out = make([]KubernetesVersionConstraint, len(*in))
		copy(*out, *in)
	}
	if in.ClassifierLabels != nil {
		in, out := &in.ClassifierLabels, &out.ClassifierLabels
		*out = make([]ClassifierLabel, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierSpec.
func (in *ClassifierSpec) DeepCopy() *ClassifierSpec {
	if in == nil {
		return nil
	}
	out := new(ClassifierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierStatus) DeepCopyInto(out *ClassifierStatus) {
	*out = *in
	if in.MachingClusterStatuses != nil {
		in, out := &in.MachingClusterStatuses, &out.MachingClusterStatuses
		*out = make([]MachingClusterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = make([]ClusterInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierStatus.
func (in *ClassifierStatus) DeepCopy() *ClassifierStatus {
	if in == nil {
		return nil
	}
	out := new(ClassifierStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.Cluster = in.Cluster
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfiguration) DeepCopyInto(out *ComponentConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfiguration.
func (in *ComponentConfiguration) DeepCopy() *ComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(ComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfiguration) DeepCopyInto(out *DebuggingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfiguration.
func (in *DebuggingConfiguration) DeepCopy() *DebuggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DebuggingConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfigurationList) DeepCopyInto(out *DebuggingConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DebuggingConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfigurationList.
func (in *DebuggingConfigurationList) DeepCopy() *DebuggingConfigurationList {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DebuggingConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfigurationSpec) DeepCopyInto(out *DebuggingConfigurationSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ComponentConfiguration, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfigurationSpec.
func (in *DebuggingConfigurationSpec) DeepCopy() *DebuggingConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployedResource) DeepCopyInto(out *DeployedResource) {
	*out = *in
	if in.LabelFilters != nil {
		in, out := &in.LabelFilters, &out.LabelFilters
		*out = make([]LabelFilter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployedResource.
func (in *DeployedResource) DeepCopy() *DeployedResource {
	if in == nil {
		return nil
	}
	out := new(DeployedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesVersionConstraint) DeepCopyInto(out *KubernetesVersionConstraint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesVersionConstraint.
func (in *KubernetesVersionConstraint) DeepCopy() *KubernetesVersionConstraint {
	if in == nil {
		return nil
	}
	out := new(KubernetesVersionConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelFilter) DeepCopyInto(out *LabelFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelFilter.
func (in *LabelFilter) DeepCopy() *LabelFilter {
	if in == nil {
		return nil
	}
	out := new(LabelFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachingClusterStatus) DeepCopyInto(out *MachingClusterStatus) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.ManagedLabels != nil {
		in, out := &in.ManagedLabels, &out.ManagedLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UnManagedLabels != nil {
		in, out := &in.UnManagedLabels, &out.UnManagedLabels
		*out = make([]UnManagedLabel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachingClusterStatus.
func (in *MachingClusterStatus) DeepCopy() *MachingClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MachingClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRef) DeepCopyInto(out *PolicyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRef.
func (in *PolicyRef) DeepCopy() *PolicyRef {
	if in == nil {
		return nil
	}
	out := new(PolicyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnManagedLabel) DeepCopyInto(out *UnManagedLabel) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnManagedLabel.
func (in *UnManagedLabel) DeepCopy() *UnManagedLabel {
	if in == nil {
		return nil
	}
	out := new(UnManagedLabel)
	in.DeepCopyInto(out)
	return out
}
