//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadLetterConfigObservation) DeepCopyInto(out *DeadLetterConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadLetterConfigObservation.
func (in *DeadLetterConfigObservation) DeepCopy() *DeadLetterConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DeadLetterConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadLetterConfigParameters) DeepCopyInto(out *DeadLetterConfigParameters) {
	*out = *in
	if in.TargetArn != nil {
		in, out := &in.TargetArn, &out.TargetArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadLetterConfigParameters.
func (in *DeadLetterConfigParameters) DeepCopy() *DeadLetterConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DeadLetterConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationConfigObservation) DeepCopyInto(out *DestinationConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationConfigObservation.
func (in *DestinationConfigObservation) DeepCopy() *DestinationConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DestinationConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationConfigParameters) DeepCopyInto(out *DestinationConfigParameters) {
	*out = *in
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = make([]OnFailureParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationConfigParameters.
func (in *DestinationConfigParameters) DeepCopy() *DestinationConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentObservation) DeepCopyInto(out *EnvironmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentObservation.
func (in *EnvironmentObservation) DeepCopy() *EnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentParameters) DeepCopyInto(out *EnvironmentParameters) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentParameters.
func (in *EnvironmentParameters) DeepCopy() *EnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralStorageObservation) DeepCopyInto(out *EphemeralStorageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralStorageObservation.
func (in *EphemeralStorageObservation) DeepCopy() *EphemeralStorageObservation {
	if in == nil {
		return nil
	}
	out := new(EphemeralStorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralStorageParameters) DeepCopyInto(out *EphemeralStorageParameters) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralStorageParameters.
func (in *EphemeralStorageParameters) DeepCopy() *EphemeralStorageParameters {
	if in == nil {
		return nil
	}
	out := new(EphemeralStorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMapping) DeepCopyInto(out *EventSourceMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMapping.
func (in *EventSourceMapping) DeepCopy() *EventSourceMapping {
	if in == nil {
		return nil
	}
	out := new(EventSourceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingList) DeepCopyInto(out *EventSourceMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSourceMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingList.
func (in *EventSourceMappingList) DeepCopy() *EventSourceMappingList {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingObservation) DeepCopyInto(out *EventSourceMappingObservation) {
	*out = *in
	if in.FunctionArn != nil {
		in, out := &in.FunctionArn, &out.FunctionArn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastProcessingResult != nil {
		in, out := &in.LastProcessingResult, &out.LastProcessingResult
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateTransitionReason != nil {
		in, out := &in.StateTransitionReason, &out.StateTransitionReason
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingObservation.
func (in *EventSourceMappingObservation) DeepCopy() *EventSourceMappingObservation {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingParameters) DeepCopyInto(out *EventSourceMappingParameters) {
	*out = *in
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(float64)
		**out = **in
	}
	if in.BisectBatchOnFunctionError != nil {
		in, out := &in.BisectBatchOnFunctionError, &out.BisectBatchOnFunctionError
		*out = new(bool)
		**out = **in
	}
	if in.DestinationConfig != nil {
		in, out := &in.DestinationConfig, &out.DestinationConfig
		*out = make([]DestinationConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventSourceArn != nil {
		in, out := &in.EventSourceArn, &out.EventSourceArn
		*out = new(string)
		**out = **in
	}
	if in.FilterCriteria != nil {
		in, out := &in.FilterCriteria, &out.FilterCriteria
		*out = make([]FilterCriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.FunctionResponseTypes != nil {
		in, out := &in.FunctionResponseTypes, &out.FunctionResponseTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaximumBatchingWindowInSeconds != nil {
		in, out := &in.MaximumBatchingWindowInSeconds, &out.MaximumBatchingWindowInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.MaximumRecordAgeInSeconds != nil {
		in, out := &in.MaximumRecordAgeInSeconds, &out.MaximumRecordAgeInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.MaximumRetryAttempts != nil {
		in, out := &in.MaximumRetryAttempts, &out.MaximumRetryAttempts
		*out = new(float64)
		**out = **in
	}
	if in.ParallelizationFactor != nil {
		in, out := &in.ParallelizationFactor, &out.ParallelizationFactor
		*out = new(float64)
		**out = **in
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SelfManagedEventSource != nil {
		in, out := &in.SelfManagedEventSource, &out.SelfManagedEventSource
		*out = make([]SelfManagedEventSourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceAccessConfiguration != nil {
		in, out := &in.SourceAccessConfiguration, &out.SourceAccessConfiguration
		*out = make([]SourceAccessConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartingPosition != nil {
		in, out := &in.StartingPosition, &out.StartingPosition
		*out = new(string)
		**out = **in
	}
	if in.StartingPositionTimestamp != nil {
		in, out := &in.StartingPositionTimestamp, &out.StartingPositionTimestamp
		*out = new(string)
		**out = **in
	}
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TumblingWindowInSeconds != nil {
		in, out := &in.TumblingWindowInSeconds, &out.TumblingWindowInSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingParameters.
func (in *EventSourceMappingParameters) DeepCopy() *EventSourceMappingParameters {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingSpec) DeepCopyInto(out *EventSourceMappingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingSpec.
func (in *EventSourceMappingSpec) DeepCopy() *EventSourceMappingSpec {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingStatus) DeepCopyInto(out *EventSourceMappingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingStatus.
func (in *EventSourceMappingStatus) DeepCopy() *EventSourceMappingStatus {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemConfigObservation) DeepCopyInto(out *FileSystemConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemConfigObservation.
func (in *FileSystemConfigObservation) DeepCopy() *FileSystemConfigObservation {
	if in == nil {
		return nil
	}
	out := new(FileSystemConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemConfigParameters) DeepCopyInto(out *FileSystemConfigParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.LocalMountPath != nil {
		in, out := &in.LocalMountPath, &out.LocalMountPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemConfigParameters.
func (in *FileSystemConfigParameters) DeepCopy() *FileSystemConfigParameters {
	if in == nil {
		return nil
	}
	out := new(FileSystemConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterCriteriaObservation) DeepCopyInto(out *FilterCriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterCriteriaObservation.
func (in *FilterCriteriaObservation) DeepCopy() *FilterCriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(FilterCriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterCriteriaParameters) DeepCopyInto(out *FilterCriteriaParameters) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterCriteriaParameters.
func (in *FilterCriteriaParameters) DeepCopy() *FilterCriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(FilterCriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterObservation) DeepCopyInto(out *FilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterObservation.
func (in *FilterObservation) DeepCopy() *FilterObservation {
	if in == nil {
		return nil
	}
	out := new(FilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterParameters) DeepCopyInto(out *FilterParameters) {
	*out = *in
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterParameters.
func (in *FilterParameters) DeepCopy() *FilterParameters {
	if in == nil {
		return nil
	}
	out := new(FilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InvokeArn != nil {
		in, out := &in.InvokeArn, &out.InvokeArn
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.QualifiedArn != nil {
		in, out := &in.QualifiedArn, &out.QualifiedArn
		*out = new(string)
		**out = **in
	}
	if in.SigningJobArn != nil {
		in, out := &in.SigningJobArn, &out.SigningJobArn
		*out = new(string)
		**out = **in
	}
	if in.SigningProfileVersionArn != nil {
		in, out := &in.SigningProfileVersionArn, &out.SigningProfileVersionArn
		*out = new(string)
		**out = **in
	}
	if in.SourceCodeSize != nil {
		in, out := &in.SourceCodeSize, &out.SourceCodeSize
		*out = new(float64)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.Architectures != nil {
		in, out := &in.Architectures, &out.Architectures
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CodeSigningConfigArn != nil {
		in, out := &in.CodeSigningConfigArn, &out.CodeSigningConfigArn
		*out = new(string)
		**out = **in
	}
	if in.DeadLetterConfig != nil {
		in, out := &in.DeadLetterConfig, &out.DeadLetterConfig
		*out = make([]DeadLetterConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvironmentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EphemeralStorage != nil {
		in, out := &in.EphemeralStorage, &out.EphemeralStorage
		*out = make([]EphemeralStorageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FileSystemConfig != nil {
		in, out := &in.FileSystemConfig, &out.FileSystemConfig
		*out = make([]FileSystemConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.ImageConfig != nil {
		in, out := &in.ImageConfig, &out.ImageConfig
		*out = make([]ImageConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArn != nil {
		in, out := &in.KMSKeyArn, &out.KMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArnRef != nil {
		in, out := &in.KMSKeyArnRef, &out.KMSKeyArnRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KMSKeyArnSelector != nil {
		in, out := &in.KMSKeyArnSelector, &out.KMSKeyArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		*out = new(float64)
		**out = **in
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.Publish != nil {
		in, out := &in.Publish, &out.Publish
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReservedConcurrentExecutions != nil {
		in, out := &in.ReservedConcurrentExecutions, &out.ReservedConcurrentExecutions
		*out = new(float64)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3Key != nil {
		in, out := &in.S3Key, &out.S3Key
		*out = new(string)
		**out = **in
	}
	if in.S3ObjectVersion != nil {
		in, out := &in.S3ObjectVersion, &out.S3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.SourceCodeHash != nil {
		in, out := &in.SourceCodeHash, &out.SourceCodeHash
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.TracingConfig != nil {
		in, out := &in.TracingConfig, &out.TracingConfig
		*out = make([]TracingConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VPCConfig != nil {
		in, out := &in.VPCConfig, &out.VPCConfig
		*out = make([]VPCConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigObservation) DeepCopyInto(out *ImageConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigObservation.
func (in *ImageConfigObservation) DeepCopy() *ImageConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ImageConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigParameters) DeepCopyInto(out *ImageConfigParameters) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EntryPoint != nil {
		in, out := &in.EntryPoint, &out.EntryPoint
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WorkingDirectory != nil {
		in, out := &in.WorkingDirectory, &out.WorkingDirectory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigParameters.
func (in *ImageConfigParameters) DeepCopy() *ImageConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ImageConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnFailureObservation) DeepCopyInto(out *OnFailureObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnFailureObservation.
func (in *OnFailureObservation) DeepCopy() *OnFailureObservation {
	if in == nil {
		return nil
	}
	out := new(OnFailureObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnFailureParameters) DeepCopyInto(out *OnFailureParameters) {
	*out = *in
	if in.DestinationArn != nil {
		in, out := &in.DestinationArn, &out.DestinationArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnFailureParameters.
func (in *OnFailureParameters) DeepCopy() *OnFailureParameters {
	if in == nil {
		return nil
	}
	out := new(OnFailureParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfManagedEventSourceObservation) DeepCopyInto(out *SelfManagedEventSourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfManagedEventSourceObservation.
func (in *SelfManagedEventSourceObservation) DeepCopy() *SelfManagedEventSourceObservation {
	if in == nil {
		return nil
	}
	out := new(SelfManagedEventSourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfManagedEventSourceParameters) DeepCopyInto(out *SelfManagedEventSourceParameters) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfManagedEventSourceParameters.
func (in *SelfManagedEventSourceParameters) DeepCopy() *SelfManagedEventSourceParameters {
	if in == nil {
		return nil
	}
	out := new(SelfManagedEventSourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceAccessConfigurationObservation) DeepCopyInto(out *SourceAccessConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceAccessConfigurationObservation.
func (in *SourceAccessConfigurationObservation) DeepCopy() *SourceAccessConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SourceAccessConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceAccessConfigurationParameters) DeepCopyInto(out *SourceAccessConfigurationParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceAccessConfigurationParameters.
func (in *SourceAccessConfigurationParameters) DeepCopy() *SourceAccessConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SourceAccessConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigObservation) DeepCopyInto(out *TracingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigObservation.
func (in *TracingConfigObservation) DeepCopy() *TracingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(TracingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigParameters) DeepCopyInto(out *TracingConfigParameters) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigParameters.
func (in *TracingConfigParameters) DeepCopy() *TracingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(TracingConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigObservation) DeepCopyInto(out *VPCConfigObservation) {
	*out = *in
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigObservation.
func (in *VPCConfigObservation) DeepCopy() *VPCConfigObservation {
	if in == nil {
		return nil
	}
	out := new(VPCConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigParameters) DeepCopyInto(out *VPCConfigParameters) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigParameters.
func (in *VPCConfigParameters) DeepCopy() *VPCConfigParameters {
	if in == nil {
		return nil
	}
	out := new(VPCConfigParameters)
	in.DeepCopyInto(out)
	return out
}
