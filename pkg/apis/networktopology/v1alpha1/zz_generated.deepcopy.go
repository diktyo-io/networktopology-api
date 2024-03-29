//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostInfo) DeepCopyInto(out *CostInfo) {
	*out = *in
	out.BandwidthCapacity = in.BandwidthCapacity.DeepCopy()
	out.BandwidthAllocated = in.BandwidthAllocated.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostInfo.
func (in *CostInfo) DeepCopy() *CostInfo {
	if in == nil {
		return nil
	}
	out := new(CostInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CostList) DeepCopyInto(out *CostList) {
	{
		in := &in
		*out = make(CostList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostList.
func (in CostList) DeepCopy() CostList {
	if in == nil {
		return nil
	}
	out := new(CostList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkTopology) DeepCopyInto(out *NetworkTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkTopology.
func (in *NetworkTopology) DeepCopy() *NetworkTopology {
	if in == nil {
		return nil
	}
	out := new(NetworkTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkTopologyList) DeepCopyInto(out *NetworkTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkTopologyList.
func (in *NetworkTopologyList) DeepCopy() *NetworkTopologyList {
	if in == nil {
		return nil
	}
	out := new(NetworkTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkTopologySpec) DeepCopyInto(out *NetworkTopologySpec) {
	*out = *in
	if in.Weights != nil {
		in, out := &in.Weights, &out.Weights
		*out = make(WeightList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkTopologySpec.
func (in *NetworkTopologySpec) DeepCopy() *NetworkTopologySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkTopologyStatistics) DeepCopyInto(out *NetworkTopologyStatistics) {
	*out = *in
	out.MinBandwidth = in.MinBandwidth.DeepCopy()
	out.AvgBandwidth = in.AvgBandwidth.DeepCopy()
	out.MaxBandwidth = in.MaxBandwidth.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkTopologyStatistics.
func (in *NetworkTopologyStatistics) DeepCopy() *NetworkTopologyStatistics {
	if in == nil {
		return nil
	}
	out := new(NetworkTopologyStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkTopologyStatus) DeepCopyInto(out *NetworkTopologyStatus) {
	*out = *in
	in.WeightCalculationTime.DeepCopyInto(&out.WeightCalculationTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkTopologyStatus.
func (in *NetworkTopologyStatus) DeepCopy() *NetworkTopologyStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkTopologyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginInfo) DeepCopyInto(out *OriginInfo) {
	*out = *in
	in.OriginInterStatistics.DeepCopyInto(&out.OriginInterStatistics)
	in.OriginIntraStatistics.DeepCopyInto(&out.OriginIntraStatistics)
	if in.CostList != nil {
		in, out := &in.CostList, &out.CostList
		*out = make(CostList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginInfo.
func (in *OriginInfo) DeepCopy() *OriginInfo {
	if in == nil {
		return nil
	}
	out := new(OriginInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OriginList) DeepCopyInto(out *OriginList) {
	{
		in := &in
		*out = make(OriginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginList.
func (in OriginList) DeepCopy() OriginList {
	if in == nil {
		return nil
	}
	out := new(OriginList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologyInfo) DeepCopyInto(out *TopologyInfo) {
	*out = *in
	if in.OriginList != nil {
		in, out := &in.OriginList, &out.OriginList
		*out = make(OriginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyInfo.
func (in *TopologyInfo) DeepCopy() *TopologyInfo {
	if in == nil {
		return nil
	}
	out := new(TopologyInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TopologyList) DeepCopyInto(out *TopologyList) {
	{
		in := &in
		*out = make(TopologyList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyList.
func (in TopologyList) DeepCopy() TopologyList {
	if in == nil {
		return nil
	}
	out := new(TopologyList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeightInfo) DeepCopyInto(out *WeightInfo) {
	*out = *in
	if in.TopologyList != nil {
		in, out := &in.TopologyList, &out.TopologyList
		*out = make(TopologyList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightInfo.
func (in *WeightInfo) DeepCopy() *WeightInfo {
	if in == nil {
		return nil
	}
	out := new(WeightInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in WeightList) DeepCopyInto(out *WeightList) {
	{
		in := &in
		*out = make(WeightList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightList.
func (in WeightList) DeepCopy() WeightList {
	if in == nil {
		return nil
	}
	out := new(WeightList)
	in.DeepCopyInto(out)
	return *out
}
